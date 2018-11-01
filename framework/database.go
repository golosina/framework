package framework

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	*gorm.DB
}

func NewDatabase() *Database {

	dbConnString := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	log.Println("Trying to connect to DB")
	log.Println(dbConnString)

	db, err := gorm.Open(os.Getenv("DB_DRIVER"), dbConnString)
	if err != nil {
		log.Println(err)
		panic("Failed to connect to Database")
	}

	return &Database{db}
}
