package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"online-shop-api/helper"
	payload "online-shop-api/model/dto/request"
	"online-shop-api/model/dto/response"
	"online-shop-api/service"
)

type CategoryController struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *CategoryController {
	return &CategoryController{CategoryService: categoryService}
}

func (controller *CategoryController) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryRequest := payload.CategoryCreateRequest{}
	err := json.NewDecoder(request.Body).Decode(&categoryRequest)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.CreateCategory(request.Context(), categoryRequest)
	commonResponse := response.CommonResponse{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data:   categoryResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(commonResponse)
	helper.PanicIfError(err)
}

func (controller *CategoryController) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")

	categoryResponse := controller.CategoryService.GetById(request.Context(), categoryId)
	commonResponse := response.CommonResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   categoryResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(commonResponse)
	helper.PanicIfError(err)
}

func (controller *CategoryController) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponse := controller.CategoryService.GetAll(request.Context())
	commonResponse := response.CommonResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   categoryResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(commonResponse)
	helper.PanicIfError(err)
}

func (controller *CategoryController) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	updateRequest := payload.CategoryUpdateRequest{}
	err := json.NewDecoder(request.Body).Decode(&updateRequest)
	helper.PanicIfError(err)

	categoryId := params.ByName("categoryId")
	updateRequest.Id = categoryId

	categoryResponse := controller.CategoryService.UpdateCategory(request.Context(), updateRequest)
	commonResponse := response.CommonResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   categoryResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(commonResponse)
	helper.PanicIfError(err)
}

func (controller *CategoryController) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")

	controller.CategoryService.DeleteCategory(request.Context(), categoryId)
	commonResponse := response.CommonResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
	}

	writer.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(commonResponse)
	helper.PanicIfError(err)
}
