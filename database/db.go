package database

import (
	"fmt"

	//"github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	// "gorm.io/driver/sqlserver"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// github.com/mattn/go-sqlite3
//db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

var Db *gorm.DB

func InitDb() *gorm.DB { // OOP constructor
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/goprojectgroup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Error...")
		return nil
	}
	return db
}
