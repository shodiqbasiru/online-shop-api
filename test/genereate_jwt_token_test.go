package test

import (
	"fmt"
	"online-shop-api/model/domain"
	"online-shop-api/utils"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	user := domain.User{
		Id:       "userId 1",
		NoHp:     "08123567",
		Email:    "email.@gmail.com",
		Password: "password",
		Role:     domain.RoleAdmin,
	}
	token, err := utils.GenerateJwtToken(user)
	if err != nil {
		t.Fatalf("Error generating token: %v", err)
	}
	fmt.Println(token)
}
