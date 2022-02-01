package db

import (
	//"CRUD-Simples/entyties"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB
var err error

func ConnectDB() {

	// Loading enviroment variables
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbname := os.Getenv("NAME")
	dbpassword := os.Getenv("PASSWORD")

	// Database connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbname, dbpassword, dbPort)

	// Openning connection to database
	DB, err = gorm.Open(dialect, dbURI)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected to database successfully")
	}

	// Close the databse connection when the main function closes
	// defer DB.Close()

	// Make migrations to the database if they haven't been made already
	// clientes := []entyties.Cliente{}
	// clientes = entyties.Clientes
	// DB.AutoMigrate(&clientes)

}
