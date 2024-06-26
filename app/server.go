package app

import (
	"net/http"
	"online-shop-api/helper"
	"online-shop-api/middleware"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:7720",
		Handler: authMiddleware,
	}
}

func RunServer() {
	server := InitializedServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
