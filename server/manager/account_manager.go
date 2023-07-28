package manager

import (
	"github.com/J-Obog/paidoff/clock"
	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/rest"
	"github.com/J-Obog/paidoff/store"
)

type AccountManager struct {
	store store.AccountStore
	clock clock.Clock
}

func (manager *AccountManager) Update(existing *data.Account, body *rest.AccountUpdateBody) (bool, error) {
	timestamp := manager.clock.Now()

	update := data.AccountUpdate{
		Name: body.Name,
	}

	existing.Name = update.Name

	return manager.store.Update(existing.Id, update, timestamp)
}

func (manager *AccountManager) Delete(id string) (bool, error) {
	return manager.store.SetDeleted(id)
}
