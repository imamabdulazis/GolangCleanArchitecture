package auth

import (
	"errors"
	"net/http"
	"time"

	"tugasakhircoffe/TaCoffe/api/models"
	"tugasakhircoffe/TaCoffe/api/responses"
	"tugasakhircoffe/TaCoffe/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(config.SECRETKEY)
}

//Extract retrieve the token from headers
func ExtractToken(w http.ResponseWriter, r *http.Request) *jwt.Token {
	token, err := request.ParseFromRequestWithClaims(
		r,
		request.OAuth2Extractor,
		&models.Claim{},
		func(t *jwt.Token) (interface{}, error) {
			return config.SECRETKEY, nil
		},
	)

	if err != nil {
		code := http.StatusUnauthorized
		switch err.(type) {
		case *jwt.ValidationError:
			vError := err.(*jwt.ValidationError)
			switch vError.Errors {
			case jwt.ValidationErrorExpired:
				err = errors.New("Your token has expired")
				responses.ERROR(w, code, err)
				return nil
			case jwt.ValidationErrorSignatureInvalid:
				err = errors.New("The signature token is invalid")
				responses.ERROR(w, code, err)
				return nil
			default:
				responses.ERROR(w, code, err)
				return nil
			}
		}
	}
	return token
}
