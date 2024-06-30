package app

import (
	"github.com/fatih/color"
	"net/http"
	"online-shop-api/internal/helper"
	"online-shop-api/internal/middleware"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:7720",
		Handler: authMiddleware,
	}
}

func RunServer() {
	server := InitializedServer()

	text := color.New(color.FgGreen).PrintFunc()
	text("Server is running on ", server.Addr, " ...\n")

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
