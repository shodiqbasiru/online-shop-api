package app

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"online-shop-api/helper"
)

func NewServer(router *httprouter.Router) *http.Server {
	return &http.Server{
		Addr:    "localhost:7720",
		Handler: router,
	}
}

func RunServer() {
	server := InitializedServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
