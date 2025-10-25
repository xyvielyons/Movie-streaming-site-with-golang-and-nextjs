package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	controller "github.com/xyvielyons/moviestreaming/controllers"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, MagicStreamMovies")
	})

	router.GET("/movies", controller.GetMovies())
	router.GET("/movie/:imdb_id", controller.GetMovie())
	router.POST("/addmovie", controller.AddMovie())

	// err := router.Run(":8000")

	// if err != nil {
	// 	fmt.Println("Failed to start server", err)
	// }

	if err := router.Run(":8000"); err != nil {
		fmt.Println("Failed to start server", err)
	}

}
