package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

// DB name of db connection
var DB *gorm.DB

// Open to open db connection
func Open() {
	// var err error
	server := "127.0.0.1"
	port := 5432
	username := "postgres"
	database := "customerdb"
	password := "postgres"
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		server, port, username, database,
		password)
	DB, err := gorm.Open("postgres", psqlInfo)
	// if err != nil {
	// 	return err
	// }
	defer DB.Close()
	if err != nil {
		// handle err
	}

	// _ = DB
	_ = err

}

// Close to close db connection
// func Close() error {
// 	return DB.Close()
// }
