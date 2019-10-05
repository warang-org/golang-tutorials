package main

import (
	"go-gorm/database"
	"go-gorm/insertrecords"

	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func main() {
	DB, err := database.Open()
	if err != nil {
		panic(err.Error())
	}

	log.Println("Connection Established")
	DB.AutoMigrate(&insertrecords.User{}, &insertrecords.UserModel{})

	dbConn := insertrecords.DBConn{DBConn: DB}
	dbConn.InsertQuery()
	dbConn.UpdateQuery()

	log.Println("Connection Established new")
}
