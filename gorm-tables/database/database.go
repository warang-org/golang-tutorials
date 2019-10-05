package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

// Open to open db connection
func Open() (*gorm.DB, error) {

	server := "127.0.0.1"
	port := 5432
	username := "postgres"
	database := "customerdb"
	password := "postgres"
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		server, port, username, database,
		password)
	DB, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return DB, nil
}
