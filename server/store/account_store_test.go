package store

import (
	"testing"

	"github.com/J-Obog/paidoff/data"
	"github.com/stretchr/testify/assert"
)

func TestAccountStore(t *testing.T) {
	it := NewStoreIntegrationTest()

	t.Run("it inserts and gets", func(t *testing.T) {
		it.Setup()

		account := testAccount()

		err := it.AccountStore.Insert(account)
		assert.NoError(t, err)

		found, err := it.AccountStore.Get(account.Id)
		assert.NoError(t, err)
		assert.NotNil(t, found)
		assert.Equal(t, account, *found)
	})

	t.Run("it updates", func(t *testing.T) {
		it.Setup()

		account := testAccount()
		account.Email = "some-old-email@gmail.com"

		newEmail := "some-new-email@gmail.com"
		update := data.AccountUpdate{Email: newEmail}

		err := it.AccountStore.Insert(account)
		assert.NoError(t, err)

		ok, err := it.AccountStore.Update(account.Id, update, 1234)
		assert.NoError(t, err)
		assert.True(t, ok)

		found, err := it.AccountStore.Get(account.Id)
		assert.NoError(t, err)
		assert.NotNil(t, found)
		assert.Equal(t, found.Email, newEmail)
	})

	t.Run("it marks as deleted", func(t *testing.T) {
		it.Setup()

		account := testAccount()
		account.IsDeleted = false

		err := it.AccountStore.Insert(account)
		assert.NoError(t, err)

		ok, err := it.AccountStore.SetDeleted(account.Id)
		assert.NoError(t, err)
		assert.True(t, ok)

		found, err := it.AccountStore.Get(account.Id)
		assert.NoError(t, err)
		assert.NotNil(t, found)
		assert.Equal(t, found.IsDeleted, true)
	})

	t.Run("it deletes", func(t *testing.T) {
		it.Setup()

		account := testAccount()

		err := it.AccountStore.Insert(account)
		assert.NoError(t, err)

		ok, err := it.AccountStore.Delete(account.Id)
		assert.NoError(t, err)
		assert.True(t, ok)

		found, err := it.AccountStore.Get(account.Id)
		assert.NoError(t, err)
		assert.Nil(t, found)
	})
}
