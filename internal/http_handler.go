package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewHttpHandler() http.Handler {
	root := gin.Default()
	apiV1 := root.Group("/api/v1")
	apiV1.GET("/ping", handlePing)
	apiV1.POST("/hello", handleHello)
	return root
}

type PingResponse struct {
	Message string `json:"message"`
}

func handlePing(c *gin.Context) {
	c.JSON(200, &PingResponse{
		Message: "pong",
	})
}

type HelloRequest struct {
	Name string `json:"name" binding:"required"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

func handleHello(c *gin.Context) {
	var request HelloRequest
	if err := c.BindJSON(&request); err != nil {
		return
	}
	c.JSON(200, &HelloResponse{
		Message: fmt.Sprintf("Hello %s!", request.Name),
	})
}
