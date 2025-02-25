package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Read environment variables
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("PORT environment variable was not specified, defaulting to 5005")
		port = "5005"
	}

	// Initialize Gin router
	router := gin.Default()

	// Don't trust any proxies, should be used when not running behind a proxy or load balancer
	router.SetTrustedProxies([]string{})

	// Health check endpoint for Kubernetes
	router.GET("/healthz", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// Simple Hello endpoint
	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello")
	})

	// Listen on 0.0.0.0:<port> so it's accessible externally (e.g., in Kubernetes)
	router.Run("0.0.0.0:" + port)
}
