package store

import (
	"github.com/BladekTech/blaze/pkg/protocol"
)

type Store struct {
	Pairs map[string]string
	// Protected bool
	// Password  string
}

type StoreResult struct {
	Result *string
	Status protocol.Status
}

func NewStore() Store {
	return Store{
		Pairs: make(map[string]string),
		// Protected: protected,
		// Password:  password,
	}
}

func (store Store) Get(key string) StoreResult {
	if store.Exists(key) {
		result := store.Pairs[key]
		return StoreResult{
			Result: &result,
			Status: protocol.STATUS_OK,
		}
	} else {
		return StoreResult{
			Result: nil,
			Status: protocol.STATUS_NO_SUCH_KEY,
		}
	}
}

func (store Store) Set(key string, value string) StoreResult {
	if !store.Exists(key) {
		store.Pairs[key] = value
		return StoreResult{
			Result: nil,
			Status: protocol.STATUS_OK,
		}
	} else {
		return StoreResult{
			Result: nil,
			Status: protocol.STATUS_KEY_ALREADY_EXISTS,
		}
	}
}

func (store Store) Update(key string, value string) StoreResult {
	if store.Exists(key) {
		store.Pairs[key] = value
		return StoreResult{
			Result: nil,
			Status: protocol.STATUS_OK,
		}
	} else {
		return StoreResult{
			Result: nil,
			Status: protocol.STATUS_NO_SUCH_KEY,
		}
	}
}

func (store Store) Delete(key string) StoreResult {
	if store.Exists(key) {
		delete(store.Pairs, key)
		return StoreResult{
			Result: nil,
			Status: protocol.STATUS_OK,
		}
	} else {
		return StoreResult{
			Result: nil,
			Status: protocol.STATUS_NO_SUCH_KEY,
		}
	}
}

func (store Store) Clear() StoreResult {
	for k := range store.Pairs {
		delete(store.Pairs, k)
	}

	return StoreResult{
		Result: nil,
		Status: protocol.STATUS_OK,
	}
}

func (store Store) Exists(key string) bool {
	return store.Pairs[key] != ""
}
