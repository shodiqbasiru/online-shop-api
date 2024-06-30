package service

import (
	"context"
	"online-shop-api/internal/model/dto/request"
	"online-shop-api/internal/model/dto/response"
)

type AuthService interface {
	RegisterUser(ctx context.Context, request request.RegisterRequest) response.RegisterResponse
	RegisterAdmin(ctx context.Context, request request.RegisterRequest) response.RegisterResponse
	LoginUser(ctx context.Context, request request.LoginRequest) response.LoginResponse
}
