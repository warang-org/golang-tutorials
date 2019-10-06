package main

import (
	"go-packnew/database"
	"go-packnew/insertrecords"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func main() {

	log.Println("Connection Established")

	database.DB.Debug().DropTableIfExists(&insertrecords.UserModel{})
	//Drops table if already exists
	database.DB.Debug().AutoMigrate(&insertrecords.UserModel{})
	//Auto create table based on Model

	config := insertrecords.UserModel{Name: "John", Address: "New York"}
	config.InsertQuery()
	log.Println("Connection Established new")
}
