package factory

import (
	"bookstore/store"
	"sync"
)

var (
	providersMu sync.RWMutex
	providers = make(map[string]store.Store)
)

