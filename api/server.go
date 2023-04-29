package api

import (
	db "github.com/Molizane/gofinance-backend/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.SQLStore
	router *gin.Engine
}

func NewServer(store *db.SQLStore) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/user", server.createUser)
	router.GET("/user/:username", server.GetUser)
	router.GET("/user/:id", server.GetUserById)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"api has an error:": err.Error()}
}
