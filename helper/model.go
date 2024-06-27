package helper

import (
	"online-shop-api/model/domain"
	"online-shop-api/model/dto/response"
)

func ToCategoryResponse(category domain.Category) response.CategoryResponse {
	return response.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []response.CategoryResponse {
	var categoryResponses []response.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}

func ToProductResponse(product domain.Product) response.ProductResponse {
	return response.ProductResponse{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		CategoryId:  product.CategoryId,
	}
}

func ToCustomerResponse(customer domain.Customer) response.CustomerResponse {
	return response.CustomerResponse{
		Id:           customer.Id,
		CustomerName: customer.CustomerName,
		Address:      customer.Address,
		Role:         customer.User.Role,
	}
}

func ToRegisterResponse(user domain.User, customer domain.Customer) response.RegisterResponse {
	return response.RegisterResponse{
		Id:           user.Id,
		CustomerName: customer.CustomerName,
		NoHp:         user.NoHp,
		Email:        user.Email,
		Role:         user.Role,
	}
}
