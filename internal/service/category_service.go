package service

import (
	"context"
	"online-shop-api/internal/model/dto/request"
	"online-shop-api/internal/model/dto/response"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, request request.CategoryCreateRequest) response.CategoryResponse
	GetById(ctx context.Context, categoryId string) response.CategoryResponse
	GetAll(ctx context.Context) []response.CategoryResponse
	UpdateCategory(ctx context.Context, request request.CategoryUpdateRequest) response.CategoryResponse
	DeleteCategory(ctx context.Context, categoryId string)
}
