package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
)

type IHandlerProvider interface {
	GetRouter() *mux.Router
}

type HandlerProvider struct {}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func (h HandlerProvider) GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", Health).Methods("GET")
	return r
}
