package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "pg-for-go-mangosteen"
	port     = 5432
	user     = "mangosteen"
	password = "123456"
	dbname   = "mangosteen_dev"
)

func Connect() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	DB = db
}

type User struct {
	ID        int
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateTables() {
	u1 := User{Email: "fyh@qq.com"}
	err := DB.Migrator().CreateTable(&u1)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("CreateTables success")
}
func Migrate() {
}

func Crud() {
	// 创建一个 User
}

func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqlDB.Close()
}
