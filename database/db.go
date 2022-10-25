package database

import (
	"fmt"

	"github.com/risdatamamal/go-assign-2/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func StartDB() (Database, error) {
	var (
		host     string = "localhost"
		port     int    = 5432
		username string = "postgres"
		password string = "postgres"
		dbName   string = "assignment_2"

		conn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbName)
	)

	db, err := gorm.Open(postgres.Open(conn))
	if err != nil {
		fmt.Println("Error connecting to database: ", err)
		return Database{}, err
	}

	err = db.Debug().AutoMigrate(models.Order{}, models.Item{})

	if err != nil {
		fmt.Println("Error when debugging and migrating database", err)
		return Database{}, err
	}

	return Database{DB: db}, nil

}
