package app

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func NewServer(router *httprouter.Router) *http.Server {
	return &http.Server{
		Addr:    "localhost:7720",
		Handler: router,
	}
}
