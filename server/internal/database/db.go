package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dsn string) (*gorm.DB,error) {
	db,err := gorm.Open(postgres.Open(dsn),&gorm.Config{})

	if err != nil{
		return nil,err
	}

	log.Println("Connected to database")

	return db,nil
}