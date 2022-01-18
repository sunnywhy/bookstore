package server

import (
	"bookstore/store"
	"bookstore/server/middleware"
	"net/http"
	"github.com/gorilla/mux"
)

type BookStoreServer struct {
	s   store.Store
	srv *http.Server
}

func NewBookStoreServer(addr string, s store.Store) *BookStoreServer {
	srv := &BookStoreServer{
		s: s,
		srv: &http.Server{
			Addr: addr,
		},
	}

	router := mux.NewRouter()
	router.HandleFunc("/book", srv.createBookHandler).Methods("POST")

	srv.srv.Handler = middleware.Logging(middleware.Validating(router))
	return srv
}