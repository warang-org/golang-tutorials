package insertrecords

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

// DBConn - DB connection
type DBConn struct {
	DBConn *gorm.DB
}

// UserModel - comments
type UserModel struct {
	gorm.Model
	Name    string `gorm:"size:255"`
	Address string `gorm:"type:varchar(100)”`
}

// User - comments
type User struct {
	gorm.Model
	Name  string
	Email string
	Age   int
}

// InsertQuery - comments for inserting records
func (db DBConn) InsertQuery() {
	user := &UserModel{Name: "John", Address: "New York"}
	usersnew := &User{Name: "Adam", Email: "adam@gmail.com", Age: 88}
	db.DBConn.Create(user)
	db.DBConn.Create(usersnew)
	// Internally it will create the query like
	// INSERT INTO 'user_models' ('name','address') VALUES ('John','New York')

	//You can insert multiple records too
	var users = []UserModel{
		UserModel{Name: "Ricky", Address: "Sydney"},
		UserModel{Name: "Adam", Address: "Brisbane"},
		UserModel{Name: "Justin", Address: "California"},
	}

	for _, user := range users {
		db.DBConn.Create(&user)
	}
	log.Println("Inserted records")
}

// UpdateQuery - update query for updating records
func (db DBConn) UpdateQuery() {
	usernew := &UserModel{Name: "John", Address: "New York"}

	users := &User{Name: "Adam", Email: "adam@gmail.com", Age: 88}

	// Select, edit, and save
	db.DBConn.Find(&usernew)
	usernew.Address = "Brisbane"
	db.DBConn.Save(&usernew)
	db.DBConn.Find(&users)
	users.Name = "JohnJohn"
	db.DBConn.Save(&users)
	// Update with column names, not attribute names
	db.DBConn.Model(&usernew).Update("Name", "Jack")
	db.DBConn.Model(&users).Update("Name", "JohnJ")
	db.DBConn.Model(&users).Update("Email", "john@gmail.com")

	db.DBConn.Model(&usernew).Updates(
		map[string]interface{}{
			"Name":    "Amy",
			"Address": "Boston",
		})
	db.DBConn.Model(&users).Updates(
		map[string]interface{}{
			"Name":  "Amy",
			"Email": "amy@gmail.com",
			"Age":   88,
		})
	// UpdateColumn()
	db.DBConn.Model(&usernew).UpdateColumn("Address", "Phoenix")
	db.DBConn.Model(&usernew).UpdateColumns(
		map[string]interface{}{
			"Name":    "Taylor",
			"Address": "Houston",
		})
	db.DBConn.Model(&users).Update("age", gorm.Expr("age * ? + ?", 2, 100))

	// Using Find()
	db.DBConn.Find(&usernew).Update("Address", "San Diego")

	// Batch Update
	db.DBConn.Table("user_models").Where("Address = ?", "California").Update("name", "Walker")
	log.Println("Updated address")

}
