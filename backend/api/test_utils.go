package api

import (
	"net/http"
	"reflect"

	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/config"
	"github.com/J-Obog/paidoff/manager"
	"github.com/J-Obog/paidoff/queue"
	"github.com/J-Obog/paidoff/rest"
	"github.com/J-Obog/paidoff/store"
	uuid "github.com/J-Obog/paidoff/uuidgen"
	"github.com/stretchr/testify/suite"
)

type ApiTestSuite struct {
	suite.Suite
	accountStore       store.AccountStore
	budgetStore        store.BudgetStore
	categoryStore      store.CategoryStore
	transactionStore   store.TransactionStore
	queue              queue.Queue
	accountManager     *manager.AccountManager
	budgetManager      *manager.BudgetManager
	categoryManager    *manager.CategoryManager
	transactionManager *manager.TransactionManager
}

func (s *ApiTestSuite) initDependencies() {
	cfg := config.Get()
	clock := clock.NewClock(cfg)
	uuidProvider := uuid.NewUuidProvider(cfg)

	storeSvc := store.NewStoreService(cfg)
	s.accountStore = storeSvc.AccountStore
	s.budgetStore = storeSvc.BudgetStore
	s.categoryStore = storeSvc.CategoryStore
	s.transactionStore = storeSvc.TransactionStore
	s.queue = queue.NewQueue(cfg)

	s.accountManager = manager.NewAccountManager(s.accountStore, clock)
	s.budgetManager = manager.NewBudgetManager(s.budgetStore, uuidProvider, clock)
	s.categoryManager = manager.NewCategoryManager(s.categoryStore, uuidProvider, clock, s.queue)
	s.transactionManager = manager.NewTransactionManager(s.transactionStore, uuidProvider, clock)
}

func (s *ApiTestSuite) OkResponse(res *rest.Response, expectedShema any) {
	s.Equal(reflect.TypeOf(res.Data), reflect.TypeOf(expectedShema))
	s.Equal(res.Status, http.StatusOK)
}

func (s *ApiTestSuite) ErrRepsonse(res *rest.Response, expectedError *rest.RestError) {
	s.Equal(res.Data, expectedError)
	s.Equal(res.Status, expectedError.Status)
}

func (s *ApiTestSuite) getJSONBody(obj any) rest.JSONBody {
	var jsonb rest.JSONBody

	err := jsonb.From(obj)
	s.NoError(err)

	return jsonb
}