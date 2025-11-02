package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nurkenti/furnitureShop/db/sqlc"
)

type Server struct {
	store  sqlc.Querier // запросы из сервера
	router *gin.Engine
}

func NewServer(store sqlc.Querier) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/users", server.CreateUser)
	router.GET("/users/email/:email", server.getUser)
	router.GET("/users/id/:id", server.getUserID)
	router.GET("/users", server.listUsers)
	router.DELETE("/users/delete/:email", server.deleteUser)
	router.POST("/users/numbers", server.numb)

	server.router = router
	return server
}

// Start runs Http server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
