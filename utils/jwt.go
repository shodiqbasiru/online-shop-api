package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"online-shop-api/internal/config"
	"online-shop-api/internal/model/domain"
	"time"
)

type JWT struct {
	Config *config.Config
}

func NewJWT(config *config.Config) *JWT {
	return &JWT{Config: config}
}

func (j *JWT) GenerateJwtToken(user domain.User) (string, error) {
	secretKey := j.Config.GetJwtConfig()

	claims := jwt.MapClaims{}
	claims["user_id"] = user.Id
	claims["role"] = user.Role
	claims["aud"] = "online-shop-api"
	claims["iss"] = "online-shop-api"
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString(secretKey)
}

func (j *JWT) VerifyJwtToken(token string) (jwt.MapClaims, error) {
	secretKey := j.Config.GetJwtConfig()

	tokenParse, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := tokenParse.Claims.(jwt.MapClaims); ok && tokenParse.Valid {
		return claims, nil
	}

	return nil, err
}
