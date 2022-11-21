package database

import (
	"duomly.com/go-bank-backend/helpers"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

// Create global variable
var DB *gorm.DB

// Create InitDatabase function

func InitDatabase() {

	database, err := gorm.Open("postgres", "postgres://gjnvspgzwlirbu:82ac86105c8dfdb5169e99871b557d4afd24eced8193bc4610a97670b814f4c3@ec2-52-23-131-232.compute-1.amazonaws.com:5432/d1vofcfoc2flv5?sslmode=require")

	helpers.HandleErr(err)

	// Set up connection pool

	database.DB().SetMaxIdleConns(20)

	database.DB().SetMaxOpenConns(200)

	DB = database

}
