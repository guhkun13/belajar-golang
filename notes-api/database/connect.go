package database

import (
	"fmt"
	"log"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/guhkun13/belajar-golang/notes-api/config"
	"github.com/guhkun13/belajar-golang/notes-api/internal/model"
)

var DB *gorm.DB


func ConnectDB() {
	var err error
	
	p := config.Config("DB_PORT")
	
	port, err := strconv.ParseInt(p, 10, 32)
	
	if err != nil {
		log.Println("Idiot")
	}
	
	dbhost := config.Config("DB_HOST")
	// dbport := config.Config("DB_PORT")
	dbname := config.Config("DB_NAME")
	dbuser := config.Config("DB_USER")
	dbpassword := config.Config("DB_PASSWORD")
	
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbhost, port, dbuser, dbpassword, dbname )
	
	DB, err = gorm.Open(postgres.Open(dsn))
	
	if err != nil {
		panic("Failed connect to database")
	}
	
	fmt.Println("Connection opened to database")
	
	// Migrate the DB
	DB.AutoMigrate(&model.Note{})
	fmt.Println("Database migrated")
	
}