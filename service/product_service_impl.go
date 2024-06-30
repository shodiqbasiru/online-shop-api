package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"online-shop-api/exception"
	"online-shop-api/helper"
	"online-shop-api/model/domain"
	"online-shop-api/model/dto/request"
	"online-shop-api/model/dto/response"
	"online-shop-api/repository"
)

type ProductServiceImpl struct {
	repository.ProductRepository
	CategoryService
	DB       *sql.DB
	Validate *validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, categoryService CategoryService, DB *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{ProductRepository: productRepository, CategoryService: categoryService, DB: DB, Validate: validate}
}

func (service *ProductServiceImpl) CreateProduct(ctx context.Context, request request.ProductCreateRequest) response.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)
	category := service.CategoryService.GetById(ctx, request.CategoryId)

	product := domain.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Stock:       request.Stock,
		CategoryId:  category.Id,
	}

	product = service.ProductRepository.Save(ctx, tx, product)
	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) GetById(ctx context.Context, productId string) response.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) GetAll(ctx context.Context) []response.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products := service.ProductRepository.FindAll(ctx, tx)

	var productResponses []response.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, helper.ToProductResponse(product))
	}

	return productResponses
}

func (service *ProductServiceImpl) UpdateProduct(ctx context.Context, request request.ProductUpdateRequest) response.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category := service.CategoryService.GetById(ctx, request.CategoryId)

	product.Name = request.Name
	product.Description = request.Description
	product.Price = request.Price
	product.Stock = request.Stock
	product.CategoryId = category.Id

	product = service.ProductRepository.Update(ctx, tx, product)

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) DeleteProduct(ctx context.Context, productId string) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.ProductRepository.Delete(ctx, tx, product)
}

func (service *ProductServiceImpl) UpdateStock(ctx context.Context, request domain.Product) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	product.Stock = request.Stock
	service.ProductRepository.Update(ctx, tx, product)
}
