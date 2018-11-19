package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name string
	Age  uint
}

func main() {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/users?charset=utf8&parseTime=True")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Automigrate the schema
	// Here gorm uses reflection to get to the struct and field names and then maps them to table and column names using
	// either the tags or the described convention
	db.AutoMigrate(&User{})

	// Create an entity for the purpose of this test
	db.Create(&User{Name: "James Bond", Age: 40})

	// Read
	// Find an user with id 1
	var user User
	db.First(&user, 1)
	fmt.Println(user)

	// Find an user with a specific name : James Bond
	db.First(&user, "Name = ?", "James Bond")
	fmt.Println(user)

	// Update - update Bond's age
	db.Model(&user).Update("Age", 41)
	fmt.Println(user)

	// Delete - delete user
	db.Delete(&user)

	// Transactions demo
	// Transactions allows the application code to assume atomicity from the database over multiple DB operations
	createTwoUsers(db)
}

func createTwoUsers(db *gorm.DB) {
	userA := User{Name: "UserA", Age: 20}
	userB := User{Name: "UserB", Age: 20}

	// Starts a trasaction
	// Tx is the handle which identifies the newly started transaction
	tx := db.Begin()
	if err := tx.Create(&userA).Error; err != nil {
		// Rollback causes the transaction to abort and bring the database to the consistent state before the
		// transaction was initiated
		tx.Rollback()
	}
	if err := tx.Create(&userB).Error; err != nil {
		tx.Rollback()
	}

	// Commit the transaction
	tx.Commit()
}
