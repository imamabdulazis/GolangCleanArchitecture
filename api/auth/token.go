package auth

import (
	"time"
	"tugasakhircoffe/TaCoffe/api/models"
	"tugasakhircoffe/TaCoffe/config"

	"github.com/dgrijalva/jwt-go"
)

//GenerateJWT creates a new token to client
func GenerateJWT(user models.User) (string, error) {
	claim := models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "Imam Abdul Azis",
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claim)
	return token.SignedString(config.SECRETKEY)
}
