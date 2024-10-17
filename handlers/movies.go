package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/logeshwarann-dev/golang-rest-apis/db"
)

var movies = []db.Movie{}

func GetAllMovies(context *gin.Context) {
	//movies, _ = ReadFile()
	movies, _ = db.ReadAllMovies()
	context.JSON(http.StatusOK, gin.H{"Movies": movies})
}

func GetMovie(context *gin.Context) {

	existingMovie, recordErr := db.ReadMovie(context.Param("id"))
	if recordErr != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": recordErr.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Record exists", "movie": existingMovie})

}

func AddMovie(context *gin.Context) {

	var newMovies []db.Movie
	var singleMovie db.Movie
	body, err := io.ReadAll(context.Request.Body)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read request body"})
		return
	}

	if err = json.Unmarshal(body, &newMovies); err != nil {
		log.Println("Request body doesn't contain a Json Array. So, Parsing Json Object.")
		if err = json.Unmarshal(body, &singleMovie); err != nil {
			log.Println(err)
			context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to Bind Request! Invalid Format"})
			return
		}
		log.Println("Request Body parsed successfully!")
		newMovies = append(newMovies, singleMovie)
	}

	for _, newMovie := range newMovies {
		err = db.CreateMovie(&newMovie)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to insert movie record in DB"})
			return
		} else {
			if newMovies[0].ID == 0 {
				newMovies = nil
			}
			newMovies = append(newMovies, newMovie)
		}

	}

	context.JSON(http.StatusCreated, gin.H{"message": "New Movie is added successfully!", "movie": newMovies})

}

func EditMovie(context *gin.Context) {

	var movieWithChanges db.Movie

	err := context.ShouldBindJSON(&movieWithChanges)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to insert movie record in DB"})
		return
	}

	existingMovie, recordErr := db.ReadMovie(context.Param("id"))
	if recordErr != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": recordErr.Error()})
		return
	}

	if err = db.UpdateMovie(existingMovie, movieWithChanges); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "New movie was updated successfully!"})

}

func EraseMovie(context *gin.Context) {

	existingMovie, recordErr := db.ReadMovie(context.Param("id"))
	if recordErr != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": recordErr.Error()})
		return
	}

	if err := db.DeletMovie(existingMovie); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Movie was deleted successfully!"})
}
