package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/logeshwarann-dev/golang-rest-apis/db"
	"github.com/logeshwarann-dev/golang-rest-apis/handlers"
)

func main() {

	server := gin.Default()

	err := db.ConnectToDB()

	if err != nil {
		log.Fatal("DB connection Failed!")
	}

	server.GET("/movies", handlers.GetAllMovies)

	server.GET("/movies/:id", handlers.GetMovie)

	server.POST("/movies", handlers.AddMovie)

	server.PUT("/movies/:id", handlers.EditMovie)

	server.DELETE("/movies/:id", handlers.EraseMovie)

	server.Run(":8080")

}
