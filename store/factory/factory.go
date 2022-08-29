package factory

import (
	"bookstore/store"
	"fmt"
	"sync"
)

var (
	providersMu sync.Mutex
	providers   = make(map[string]store.Store)
)

func Register(name string, p store.Store) {
	providersMu.Lock()
	defer providersMu.Unlock()
	if p == nil {
		panic("store: Register provider is nil")
	}

	if _, dup := providers[name]; dup {
		panic("store: Register called twice for provider " + name)
	}

	providers[name] = p
}

func New(providername string) (store.Store, error) {
	providersMu.RLock()
	p, ok := providers[providerName]
	providersMu.RUnlock()

	if !ok {
		return nil, fmt.Errorf("store: unknown providers %s", providerName)
	}

	return p, nil
}
