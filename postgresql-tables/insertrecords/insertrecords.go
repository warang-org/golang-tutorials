package insertrecords

import (
	"go-packnew/database"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

// UserModel - comments
type UserModel struct {
	ID      int    `gorm:"primary_key";"AUTO_INCREMENT"`
	Name    string `gorm:"size:255"`
	Address string `gorm:"type:varchar(100)”`
}

// Model definition given by gorm
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// User - comments
type User struct {
	gorm.Model // fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`will be added
	Name       string
}

// InsertQuery - comments for inserting records
func (usercheck UserModel) InsertQuery() {
	user := &UserModel{Name: "", Address: ""}
	database.DB.Create(user)

	// Internally it will create the query like
	// INSERT INTO 'user_models' ('name','address') VALUES ('John','New York')

	//You can insert multiple records too
	var users = []UserModel{
		UserModel{Name: "Ricky", Address: "Sydney"},
		UserModel{Name: "Adam", Address: "Brisbane"},
		UserModel{Name: "Justin", Address: "California"},
	}

	for _, user := range users {
		database.DB.Create(&user)
		log.Println("Inserted records")
	}
}
func updatequery() {
	usernew := &UserModel{Name: "John", Address: "New York"}
	// Select, edit, and save
	database.DB.Find(&usernew)
	usernew.Address = "Brisbane"
	database.DB.Save(&usernew)

	// Update with column names, not attribute names
	database.DB.Model(&usernew).Update("Name", "Jack")

	database.DB.Model(&usernew).Updates(
		map[string]interface{}{
			"Name":    "Amy",
			"Address": "Boston",
		})

	// UpdateColumn()
	database.DB.Model(&usernew).UpdateColumn("Address", "Phoenix")
	database.DB.Model(&usernew).UpdateColumns(
		map[string]interface{}{
			"Name":    "Taylor",
			"Address": "Houston",
		})
	// Using Find()
	database.DB.Find(&usernew).Update("Address", "San Diego")

	// Batch Update
	database.DB.Table("user_models").Where("address = ?", "california").Update("name", "Walker")
	log.Println("Updated address")
}
