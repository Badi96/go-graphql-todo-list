package main

import (
	"os"

	"github.com/badi96/go-graphql-todo-list/graph"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// main is the entry point for the application.
func main() {
	// gin setup
	gin.SetMode(gin.ReleaseMode)
	// create gin engine
	router := gin.Default()
	// enable cors
	router.Use(cors.Default())
	// create graphql endpoint
	router.POST("/todo", graph.TodoGraphRouter)
	// add webserver port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// start webserver
	router.Run(":" + port)
}
