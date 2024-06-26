package middleware

import (
	"encoding/json"
	"net/http"
	"online-shop-api/helper"
	"online-shop-api/model/dto/response"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "RAHASIA" == request.Header.Get("X-API-Key") {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		commonResponse := response.CommonResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		err := json.NewEncoder(writer).Encode(commonResponse)
		helper.PanicIfError(err)
	}
}
