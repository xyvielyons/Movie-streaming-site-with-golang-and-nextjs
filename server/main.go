package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	routes "github.com/xyvielyons/moviestreaming/routes"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, MagicStreamMovies")
	})
	routes.SetupUnProtectedRoutes(router)

	routes.SetupProtectedRoutes(router)
	// err := router.Run(":8000")

	// if err != nil {
	// 	fmt.Println("Failed to start server", err)
	// }

	if err := router.Run(":8000"); err != nil {
		fmt.Println("Failed to start server", err)
	}

}
