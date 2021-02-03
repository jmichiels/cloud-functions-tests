package internal

import (
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

func NewHttpHandler(firebaseApp *firebase.App) (http.Handler, error) {
	root := gin.Default()
	root.GET("/ping", handlePing)
	root.POST("/hello", handleHello)

	firestoreClt, err := firebaseApp.Firestore(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "get firestore client")
	}
	repo := newFirestoreRepository(firestoreClt)
	srv := newService(repo)
	rest := newRestApi(srv)

	rest.registerRoutes(root.Group("/api/v1"))
	return root, nil
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
