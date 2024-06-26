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

type AuthController struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

func (controller *AuthController) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	registerRequest := payload.RegisterRequest{}
	err := json.NewDecoder(request.Body).Decode(&registerRequest)
	helper.PanicIfError(err)

	registerResponse := controller.AuthService.RegisterUser(request.Context(), registerRequest)
	commonResponse := response.CommonResponse{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data:   registerResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(commonResponse)
	helper.PanicIfError(err)
}

func (controller *AuthController) LoginUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	loginRequest := payload.LoginRequest{}
	err := json.NewDecoder(request.Body).Decode(&loginRequest)
	helper.PanicIfError(err)

	loginResponse := controller.AuthService.LoginUser(request.Context(), loginRequest)
	commonResponse := response.CommonResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   loginResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(commonResponse)
	helper.PanicIfError(err)
}
