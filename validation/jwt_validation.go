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
}