package store

import (
	mystore "bookstore/store"
	factory "bookstore/store/factory"
	"sync"
)

func init() {
	factory.Register("mem", &MemStore{
		books: make(map[string]*mystore.Book)
	})
}

type MemStore struct {
	sync.RwMutex
	books map[string]*mystore.Book
}