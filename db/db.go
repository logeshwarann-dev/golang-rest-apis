package db

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbUser  = "postgres"
	dbPwd   = "postgres"
	dbName  = "moviesdb"
	host    = "localhost"
	port    = "5432"
	sslmode = "disable"
)

var Db *gorm.DB

func ConnectToDB() (err error) {

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", dbUser, dbPwd, dbName, host, port, sslmode)

	Db, err = gorm.Open(postgres.Open(connStr))

	if err != nil {
		return
	}

	return

}

func CreateMovie(movie *Movie) error {

	result := Db.Create(movie)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func ReadAllMovies() (movie []Movie, err error) {
	result := Db.Find(&movie)
	err = result.Error
	if err != nil {
		return nil, err
	}
	return
}

func ReadMovie(queryParam string) (Movie, error) {
	var existMovie Movie
	dbOps := Db.Where("id = ?", queryParam).First(&existMovie)
	if dbOps.Error != nil {
		log.Println("Error: ", dbOps.Error)
		return Movie{}, errors.New("no record found")
	}
	return existMovie, nil
}

func UpdateMovie(oldMovie, newMovie Movie) error {

	//update the movie with new values

	dbOps := Db.Model(&oldMovie).Updates(newMovie)
	if dbOps.Error != nil {
		log.Println("Error: ", dbOps.Error)
		return errors.New("unable to update the record")
	}
	return nil
}

func DeletMovie(eraseMovie Movie) error {

	dbOps := Db.Delete(&eraseMovie)
	if dbOps.Error != nil {
		log.Println("Error: ", dbOps.Error)
		return errors.New("unable to delete the record")
	}

	return nil

}
