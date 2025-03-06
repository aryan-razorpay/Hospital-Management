package config

import (
	"fmt"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:root@tcp(127.0.0.1:3306)/hospital?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
	fmt.Println("Database connected!")
}

// package Config

// import (
// 	"fmt"
// 	"github.com/jinzhu/gorm"
// )

// var DB *gorm.DB

// type DBConfig struct {
// 	Host     string
// 	Port     int
// 	User     string
// 	DBName   string
// 	Password string
// }

// func BuildDBConfig() *DBConfig {
// 	dbConfig := DBConfig{
// 		Host:     "localhost",
// 		Port:     3306,
// 		User:     "root",
// 		Password: "1234",
// 		DBName:   "first_go",
// 	}
// 	return &dbConfig
// }
// func DbURL(dbConfig *DBConfig) string {
// 	return fmt.Sprintf(
// 		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
// 		dbConfig.User,
// 		dbConfig.Password,
// 		dbConfig.Host,
// 		dbConfig.Port,
// 		dbConfig.DBName,
// 	)
// }
