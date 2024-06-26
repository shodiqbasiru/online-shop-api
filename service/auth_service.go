package service

import (
	"context"
	"online-shop-api/model/dto/request"
	"online-shop-api/model/dto/response"
)

type AuthService interface {
	RegisterUser(ctx context.Context, request request.RegisterRequest) response.RegisterResponse
	LoginUser(ctx context.Context, request request.LoginRequest) response.LoginResponse
}
