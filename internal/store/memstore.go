package store

import (
	mystore "bookstore/store"
	factory "bookstore/store/factory"
	"sync"
)

func init() {
	factory.Register("mem", &MemStore{
		books: make(map[string]*mystore.Book),
	})
}

type MemStore struct {
	sync.RWMutex
	books map[string]*mystore.Book
}

func (ms *MemStore) Create(book *mystore.Book) error {
	ms.Lock()
	defer ms.Unlock()

	if _, ok := ms.books[book.Id]; ok {
		return mystore.ErrExist
	}

	newBook := *book
	ms.books[book.Id] = &newBook
	return nil
}