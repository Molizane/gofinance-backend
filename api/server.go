package api

import (
	"net/http"

	db "github.com/Molizane/gofinance-backend/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.SQLStore
	Router *gin.Engine
}

func CORSConfig() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		context.Writer.Header().Set("Access-Control-Allow-Methos", "POST, DELETE, GET, PUT")

		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
			return
		}

		context.Next()
	}
}

func NewServer(store *db.SQLStore) *Server {
	server := &Server{store: store}
	router := gin.Default()
	router.Use(CORSConfig())

	router.POST("/user", server.createUser)
	router.GET("/user/:username", server.getUser)
	router.GET("/user/id/:id", server.getUserById)

	router.POST("/category", server.createCategory)
	router.GET("/category/id/:id", server.getCategory)
	router.GET("/category", server.getCategories)
	router.DELETE("/category/:id", server.deleteCategory)
	router.PUT("/category/:id", server.updateCategory)

	router.POST("/account", server.createAccount)
	router.GET("/account/id/:id", server.getAccount)
	router.GET("/account", server.getAccounts)
	router.GET("/account/graph/:user_id/:type", server.getAccountsGraph)
	router.GET("/account/reports/:user_id/:type", server.getAccountsReports)
	router.DELETE("/account/:id", server.deleteAccount)
	router.PUT("/account/:id", server.updateAccount)

	router.POST("/login", server.login)

	server.Router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}

func errorResponse(err error, extra ...string) gin.H {
	erro := err.Error()

	for _, ex := range extra {
		erro += "\n" + ex
	}

	return gin.H{"error": erro}
}

func errorResponseStr(err string) gin.H {
	return gin.H{"error": err}
}
