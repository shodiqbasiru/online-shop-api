package middleware

import (
	"encoding/json"
	"net/http"
	"online-shop-api/internal/helper"
	"online-shop-api/internal/model/dto/response"
	"online-shop-api/utils"
	"strings"
)

type AuthMiddleware struct {
	Handler http.Handler
	JWT     *utils.JWT
}

func NewAuthMiddleware(handler http.Handler, JWT *utils.JWT) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler, JWT: JWT}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	if request.URL.Path == "/api/auth/register" ||
		request.URL.Path == "/api/auth/register-admin" ||
		request.URL.Path == "/api/auth/login" {
		middleware.Handler.ServeHTTP(writer, request)
		return
	}

	authorization := request.Header.Get("Authorization")
	if authorization != "" {
		s := strings.Split(authorization, " ")
		if len(s) == 2 && s[0] == "Bearer" {
			token := s[1]
			_, err := middleware.JWT.VerifyJwtToken(token)
			if err == nil {
				middleware.Handler.ServeHTTP(writer, request)
				return
			}
		}
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusUnauthorized)

	commonResponse := response.CommonResponse{
		Code:   http.StatusUnauthorized,
		Status: "UNAUTHORIZED",
	}

	err := json.NewEncoder(writer).Encode(commonResponse)
	helper.PanicIfError(err)

}
