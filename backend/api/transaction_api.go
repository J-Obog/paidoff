package api

import (
	"github.com/J-Obog/paidoff/config"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/manager"
	"github.com/J-Obog/paidoff/rest"
)

type TransactionAPI struct {
	transactionManager *manager.TransactionManager
	categoryManager    *manager.CategoryManager
}

func NewTransactionAPI(
	transactionManager *manager.TransactionManager,
	categoryManager *manager.CategoryManager,
) *TransactionAPI {
	return &TransactionAPI{
		transactionManager: transactionManager,
		categoryManager:    categoryManager,
	}
}

func (api *TransactionAPI) Get(req *rest.Request) *rest.Response {
	id := req.Params.GetTransactionId()
	accountId := testAccountId

	transaction, err := api.transactionManager.Get(id, accountId)
	if err != nil {
		return rest.Err(err)
	}

	if transaction == nil {
		return rest.Err(rest.ErrInvalidTransactionId)
	}

	return rest.Ok(transaction)
}

func (api *TransactionAPI) GetByQuery(req *rest.Request) *rest.Response {
	var query rest.TransactionQuery
	accountId := testAccountId

	if err := req.Query.To(&query); err != nil {
		return rest.Err(err)
	}

	transactions, err := api.transactionManager.GetByQuery(accountId, query)
	if err != nil {
		return rest.Err(err)
	}

	return rest.Ok(transactions)
}

func (api *TransactionAPI) Create(req *rest.Request) *rest.Response {
	var body rest.TransactionCreateBody
	accountId := testAccountId

	if err := req.Body.To(&body); err != nil {
		return rest.Err(err)
	}

	if err := api.validateCreate(accountId, body); err != nil {
		return rest.Err(err)
	}

	transaction, err := api.transactionManager.Create(accountId, body)
	if err != nil {
		return rest.Err(err)
	}

	return rest.Ok(transaction)
}

func (api *TransactionAPI) Update(req *rest.Request) *rest.Response {
	var body rest.TransactionUpdateBody
	id := req.Params.GetTransactionId()
	accountId := testAccountId

	if err := req.Body.To(&body); err != nil {
		return rest.Err(err)
	}

	existing, err := api.transactionManager.Get(id, accountId)
	if err != nil {
		return rest.Err(err)
	}

	if err := api.validateUpdate(existing, body); err != nil {
		return rest.Err(err)
	}

	ok, err := api.transactionManager.Update(existing, body)

	if err != nil {
		return rest.Err(err)
	}

	if !ok {
		return rest.Err(rest.ErrInvalidTransactionId)
	}

	return rest.Ok(existing)
}

func (api *TransactionAPI) Delete(req *rest.Request) *rest.Response {
	id := req.Params.GetTransactionId()
	accountId := testAccountId

	ok, err := api.transactionManager.Delete(id, accountId)
	if err != nil {
		return rest.Err(err)
	}

	if !ok {
		return rest.Err(rest.ErrInvalidTransactionId)
	}

	return rest.Success()
}

func (api *TransactionAPI) validateUpdate(existing *data.Transaction, body rest.TransactionUpdateBody) error {
	if existing == nil {
		return rest.ErrInvalidTransactionId
	}

	d := data.NewDate(body.Month, body.Day, body.Year)

	if ok := d.IsValid(); !ok {
		return rest.ErrInvalidDate
	}

	if body.Note != nil {
		if err := api.checkNote(*body.Note); err != nil {
			return err
		}

	}

	if body.CategoryId != nil {
		if existing.CategoryId == nil || (*existing.CategoryId != *body.CategoryId) {
			ok, err := api.categoryManager.Exists(*body.CategoryId, existing.AccountId)

			if err != nil {
				return err
			}

			if !ok {
				return rest.ErrInvalidCategoryId
			}
		}
	}

	return nil
}

func (api *TransactionAPI) validateCreate(accountId string, body rest.TransactionCreateBody) error {
	d := data.NewDate(body.Month, body.Day, body.Year)

	if ok := d.IsValid(); !ok {
		return rest.ErrInvalidDate
	}

	if body.Note != nil {
		if err := api.checkNote(*body.Note); err != nil {
			return err
		}
	}

	if body.CategoryId != nil {
		ok, err := api.categoryManager.Exists(*body.CategoryId, accountId)

		if err != nil {
			return err
		}

		if !ok {
			return rest.ErrInvalidCategoryId
		}
	}

	return nil
}

func (api *TransactionAPI) checkNote(note string) error {
	if len(note) > config.LimitMaxTransactionNoteChars {
		return rest.ErrInvalidTransactionNote
	}

	return nil
}
