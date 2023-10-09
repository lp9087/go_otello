package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var PostgresDB *gorm.DB
var err error

func Connect(connectionString string) {
	PostgresDB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to Database...")
}

//func Migrate() {
//	err = PostgresDB.AutoMigrate(&Hotel{})
//	if err != nil {
//		log.Fatal(err)
//	}
//	log.Println("Database Migration Completed...")
//}
