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

func (ms *MemStore) Update(book *mystore.Book) error {
	ms.Lock()
	defer ms.Unlock()

	oldBook, ok := ms.books[book.Id]
	if !ok {
		return mystore.ErrNotFound
	}

	newBook := *oldBook
	if book.Name != "" {
		newBook.Name = book.Name
	}
	if book.Authors != nil {
		newBook.Authors = book.Authors
	}
	if book.Press != "" {
		newBook.Press = book.Press
	}

	ms.books[book.Id] = &newBook
	return nil
}