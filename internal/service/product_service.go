package service

import (
	"context"
	"online-shop-api/internal/model/domain"
	"online-shop-api/internal/model/dto/request"
	"online-shop-api/internal/model/dto/response"
)

type ProductService interface {
	CreateProduct(ctx context.Context, request request.ProductCreateRequest) response.ProductResponse
	GetById(ctx context.Context, productId string) response.ProductResponse
	GetAll(ctx context.Context) []response.ProductResponse
	UpdateProduct(ctx context.Context, request request.ProductUpdateRequest) response.ProductResponse
	DeleteProduct(ctx context.Context, productId string)
	UpdateStock(ctx context.Context, request domain.Product)
}
