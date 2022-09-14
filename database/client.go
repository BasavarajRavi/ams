package database

import (
	"ams/entities"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func Connect(connectionString string) {
	Instance, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Can not Connect to Database")
	}
	log.Println("Connected to Database")

}
func Migrate() {
	Instance.AutoMigrate(&entities.User{}, &entities.Employee{}, &entities.Asset{})
	log.Println("Data Migration Completed")
}
