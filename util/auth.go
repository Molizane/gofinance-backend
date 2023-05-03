package util

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func ValidateToken(ctx *gin.Context, token string) error {
	claims := &Claims{}
	var jwtSignedKey = []byte("secret_key")

	tokenParse, err := jwt.ParseWithClaims(token, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtSignedKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ctx.JSON(http.StatusUnauthorized, err.Error())
			return err
		}

		ctx.JSON(http.StatusBadRequest, err.Error())
		return err
	}

	if !tokenParse.Valid {
		ctx.JSON(http.StatusUnauthorized, "token is invalid")
		return errors.New("token is invalid")
	}

	ctx.Next()
	return nil
}

func GetTokenInHeaderAndVerify(ctx *gin.Context) error {
	authorizationHeaderKey := ctx.GetHeader("authorization")
	fields := strings.Fields(authorizationHeaderKey)

	if len(fields) < 2 {
		err := errors.New(" Please, login")
		ctx.JSON(http.StatusBadRequest, err)
		return err
	}

	tokenToValidate := fields[1]
	err := ValidateToken(ctx, tokenToValidate)

	return err
}
