package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/xyvielyons/moviestreaming/controllers"
	middleware "github.com/xyvielyons/moviestreaming/middleware"
)

func SetupProtectedRoutes(router *gin.Engine) {
	router.Use(middleware.AuthMiddleWare())

	router.GET("/movie/:imdb_id", controller.GetMovie())
	router.POST("/addmovie", controller.AddMovie())

}
