package server

import (
	"bookstore/store"
	"net/http"
)

type BookStoreServer struct {
	s   store.Store
	srv *http.Server
}
