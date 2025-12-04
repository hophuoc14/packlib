package validation

import (
	"packlib/config"
	"packlib/entity"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	Sub string `json:"sub"`
	Username string `json:"username"`
	Role string `json:"role"`
	jwt.StandardClaims
}

func CreateToken(user entity.User) (string, error) {
	configuration := config.New()
	
	claims := UserClaims{
		Sub: user.Id.String(),
		Username: user.Username,
		Role: "",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte(configuration.Get("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}