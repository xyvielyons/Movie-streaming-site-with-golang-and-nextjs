package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	database "github.com/xyvielyons/moviestreaming/database"
	models "github.com/xyvielyons/moviestreaming/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var movieCollection *mongo.Collection = database.OpenCollection("movies")

func GetMovies() gin.HandlerFunc {
	return func(c *gin.Context) {
		//This code is for memory management
		//helps us to cleanup reasources within our memory

		//this creates a new context called "ctx" that automatically cancel after a specified duration
		//ctx is the new context that carries the timeout
		//cancel helps us to manually cancel the context before the timeout if necessary
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var movies []models.Movie

		//GetMovies → Multiple documents → .Find() → use cursor.Close() and cursor.All().
		cursor, err := movieCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies"})
		}

		//Defer delays the execution of cancel until the surrounding function returns
		//this ensures the context is properly cleaned up and resources are released
		defer cursor.Close(ctx)

		if err = cursor.All(ctx, &movies); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode movies"})
		}

		c.JSON(http.StatusOK, movies)
	}
}

func GetMovie() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		movieID := c.Param("imdb_id")

		if movieID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Movie ID is required"})
			return
		}

		var movie models.Movie

		//GetMovie → Single document → .FindOne() → use .Decode() only

		err := movieCollection.FindOne(ctx, bson.M{"imdb_id": movieID}).Decode(&movie)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
			return
		}

		c.JSON(http.StatusOK, movie)

	}
}
