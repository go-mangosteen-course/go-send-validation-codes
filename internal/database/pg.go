package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"
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
}
func Migrate() {
	// 给user添加手机字段
	_, err := DB.Exec(`ALTER TABLE users ADD COLUMN phone VARCHAR(50)`)
	if err != nil {
		log.Println(err)
	}

	_, err = DB.Exec(`ALTER TABLE users ADD COLUMN address VARCHAR(200)`)
	if err != nil {
		log.Println(err)
	}

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
	if err != nil {
		log.Println(err)
	}

	_, err = DB.Exec(`ALTER TABLE items ALTER COLUMN happened_at TYPE TIMESTAMP`)
	if err != nil {
		log.Println(err)
	}

	// 给 users 的 email 字段添加唯一性索引
	_, err = DB.Exec(`CREATE UNIQUE INDEX users_email_index ON users (email)`)
	if err != nil {
		log.Println(err)
	}
}

func Crud() {
	// 创建一个 User
	_, err := DB.Exec(`INSERT INTO users (email) values ('2@qq.com')`)
	if err != nil {
		switch x := err.(type) {
		case *pq.Error:
			pqErr := err.(*pq.Error)
			log.Println(pqErr.Code.Name(), pqErr.Message)
		default:
			log.Println(x)
		}
	}
	_, err = DB.Exec(`Update users SET phone = 138123456789 where email = '2@qq.com'`)
	if err != nil {
		log.Println(err)
	}
	stmt, err := DB.Prepare("SELECT phone FROM users where email = $1 offset $2 limit $3")
	if err != nil {
		log.Fatalln(err)
	}
	result, err := stmt.Query("2@qq.com", 0, 3)
	if err != nil {
		log.Println(err)
	} else {
		for result.Next() {
			var phone string
			result.Scan(&phone)
			log.Println("phone", phone)
		}
	}

}

func PgClose() {
	DB.Close()
}
