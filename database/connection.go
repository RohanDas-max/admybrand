package database

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/rohandas-max/admybrand/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	// dsn := "rohan:abc123@tcp(127.0.0.1:3306)/assgn?charset=utf8mb4&parseTime=True&loc=Local"
	envMap, _ := godotenv.Read(".env")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata", envMap["POSTGRESQL_ADDON_HOST"], envMap["POSTGRESQL_ADDON_USER"], envMap["POSTGRESQL_ADDON_PASSWORD"], envMap["POSTGRESQL_ADDON_DB"], envMap["POSTGRESQL_ADDON_PORT"])
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		DB = db
		var model model.Data
		if res := db.AutoMigrate(&model); res != nil {
			fmt.Println("table creation error", err)
		} else {
			fmt.Println("Database Connected")
		}
	}
}
