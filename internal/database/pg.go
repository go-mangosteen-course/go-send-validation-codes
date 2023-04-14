package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "pg-for-go-mangosteen"
	port     = 5432
	user     = "mangosteen"
	password = "123456"
	dbname   = "mangosteen_dev"
)

func PgConnect() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	DB = db
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connect to db")
}

func PgCreateTables() {
	// 创建 users 表
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Successfully create users table")
}
func handleError(err error) {
	if err != nil {
		log.Println(err)
	}
}
func Migrate() {
	// 给user添加手机字段
	_, err := DB.Exec(`ALTER TABLE users ADD COLUMN phone VARCHAR(50)`)
	handleError(err)
	log.Println("Successfully add phone column to users table")

	_, err = DB.Exec(`ALTER TABLE users ADD COLUMN address VARCHAR(200)`)
	handleError(err)
	log.Println("Successfully add address column to users table")

	// 新增 Items 表，字段为 id, amount, happened_at, created_at, updated_at
	_, err = DB.Exec(`
	CREATE TABLE items(
		id SERIAL PRIMARY KEY,
		amount INTEGER NOT NULL,
		happened_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
 );
	`)
	handleError(err)
	log.Println("Successfully create items table")

	_, err = DB.Exec(`ALTER TABLE items ALTER COLUMN happened_at TYPE TIMESTAMP`)
	handleError(err)
	log.Println("Successfully change the type of happened_at to TIMESTAMP")
}

func PgClose() {
	DB.Close()
	log.Println("Successfully close db")
}
