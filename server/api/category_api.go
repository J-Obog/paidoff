package api

import (
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/manager"
	"github.com/J-Obog/paidoff/rest"
)

type CategoryAPI struct {
	categoryManager *manager.CategoryManager
}

func (api *CategoryAPI) GetCategory(r *rest.Request) *rest.Response {
	category, errRes := api.categoryCtx(r)
	if errRes != nil {
		return errRes
	}

	return buildOKResponse(category)
}

func (api *CategoryAPI) GetCategories(r *rest.Request) *rest.Response {
	categories, err := api.categoryManager.GetByAccount(r.Account.Id)

	if err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(categories)
}

func (api *CategoryAPI) CreateCategory(r *rest.Request) *rest.Response {
	reqBody, err := r.Body.CategoryCreateBody()
	if err != nil {
		return buildBadRequestError()
	}

	// check name doesnt already exist

	if err := api.categoryManager.Create(r.Account.Id, reqBody); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}

func (api *CategoryAPI) UpdateCategory(r *rest.Request) *rest.Response {
	category, errRes := api.categoryCtx(r)
	if errRes != nil {
		return errRes
	}

	reqBody, err := r.Body.CategoryUpdateBody()
	if err != nil {
		return buildBadRequestError()
	}

	// check name doesnt already exist

	if err := api.categoryManager.Update(&category, reqBody); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}

func (api *CategoryAPI) DeleteCategory(r *rest.Request) *rest.Response {
	category, errRes := api.categoryCtx(r)
	if errRes != nil {
		return errRes
	}

	if err := api.categoryManager.Delete(category.Id); err != nil {
		return buildServerError(err)
	}

	return buildOKResponse(nil)
}

func (api *CategoryAPI) categoryCtx(r *rest.Request) (data.Category, *rest.Response) {
	category, err := api.categoryManager.Get(r.Params.CategoryId())

	if err != nil {
		return data.Category{}, buildServerError(err)
	}
	if category == nil || category.AccountId != r.Account.Id {
		return data.Category{}, buildBadRequestError()
	}

	return *category, nil
}