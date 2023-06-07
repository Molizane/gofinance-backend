package api

import (
	"bytes"
	"crypto/sha512"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Username string `json:"username" binding:"required"`
	jwt.RegisteredClaims
}

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginRespone struct {
	UserID int32  `json:"user_id" binding:"required"`
	Token  string `json:"token" binding:"required"`
}

func (server *Server) login(ctx *gin.Context) {
	var req loginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err, "ctx.ShouldBindJSON"))
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusUnauthorized, errorResponseStr("Invalid username or password"))
			fmt.Print("Erro in server.store.GetUser - Invalid username or password - ", req.Username, "\n")
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err, "server.store.GetUser"))
    fmt.Print("Error in server.store.GetUser")
		return
	}

	hashedInput := sha512.Sum512_256([]byte(req.Password))
	trimmedHash := bytes.Trim(hashedInput[:], "\x00")
	preparedPassword := string(trimmedHash)
	plainTextInBytes := []byte(preparedPassword)
	hashTextInBytes := []byte(user.Password)

	err = bcrypt.CompareHashAndPassword(hashTextInBytes, plainTextInBytes)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponseStr("Invalid username or password"))
    fmt.Print("Error in bcrypt.CompareHashAndPassword\n")
		return
	}

	expirationTime := time.Now().Add(time.Hour * 24)

	claims := &Claims{
		Username: req.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	generatedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var jwtSignedKey = []byte("secret_key")
	generatedTokenToString, err := generatedToken.SignedString(jwtSignedKey)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
    fmt.Print("Error in generatedToken.SignedString")
		return
	}

	var arg = &loginRespone{
		UserID: user.ID,
		Token:  generatedTokenToString,
	}

	ctx.JSON(http.StatusOK, arg)
}
