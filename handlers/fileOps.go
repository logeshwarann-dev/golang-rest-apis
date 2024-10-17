package handlers

import (
	"encoding/json"
	"log"
	"os"

	"github.com/logeshwarann-dev/golang-rest-apis/db"
)

const filePath string = "movies.json"

func WriteToFile(movieDetails []db.Movie) error {

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	content, _ := json.Marshal(movieDetails)
	_, err = file.Write(content)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil

}

func ReadFile() (moviesFromFile []db.Movie, err error) {

	content, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
		return
	}

	if len(content) > 0 {
		json.Unmarshal(content, &moviesFromFile)
	}

	return

}
