package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"mangosteen/config/queries"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var DB *sql.DB
var DBCtx = context.Background()

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
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	DB = db
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
}

type User struct {
	ID        int
	Email     string `gorm:"uniqueIndex"`
	Phone     string
	Address2  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Item struct {
	ID         int
	UserID     int
	Amount     int
	HappenedAt time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
type Tag struct {
	ID   int
	Name string
}

func CreateMigration(filename string) {
	cmd := exec.Command("migrate", "create", "-ext", "sql", "-dir", "config/migrations", "-seq", filename)
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
func Migrate() {
	dir, err := os.Getwd()
	name := filepath.Base(dir)
	for !strings.Contains(name, "mangosteen") {
		dir = filepath.Dir(dir)
		name = filepath.Base(dir)
	}
	if err != nil {
		log.Fatalln(err)
	}

	m, err := migrate.New(
		fmt.Sprintf("file://%s/config/migrations", dir),
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			user, password, host, port, dbname,
		))
	if err != nil {
		log.Fatalln(err)
	}
	err = m.Up()
	if err != nil {
		if !strings.Contains(err.Error(), "no change") {
			log.Fatalln(err)
		}
	}
}

func MigrateDown() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	m, err := migrate.New(
		fmt.Sprintf("file://%s/config/migrations", dir),
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			user, password, host, port, dbname,
		))
	if err != nil {
		log.Fatalln(err)
	}
	err = m.Steps(-1)
	if err != nil {
		log.Fatalln(err)
	}
}

func Crud() {
	q := queries.New(DB)
	id := rand.Int()
	u, err := q.CreateUser(DBCtx, fmt.Sprintf("%d@qq.com", id))
	if err != nil {
		log.Fatalln(err)
	}
	err = q.UpdateUser(DBCtx, queries.UpdateUserParams{
		ID:      u.ID,
		Email:   u.Email,
		Phone:   u.Phone,
		Address: "中国杭州滨江网商路699号",
	})
	if err != nil {
		log.Fatalln(err)
	}

	// users, err := q.ListUsers(DBCtx, queries.ListUsersParams{
	// 	Offset: 0,
	// 	Limit:  10,
	// })
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Println(users)
	u, err = q.FindUserByEmail(DBCtx, fmt.Sprintf("%d@qq.com", id))
	if err != nil {
		log.Fatalln(err)
	}
	// log.Println(u)
	// err = q.DeleteUser(DBCtx, u.ID)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// users, err = q.ListUsers(DBCtx, queries.ListUsersParams{
	// 	Offset: 0,
	// 	Limit:  10,
	// })
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Println(users)
}

func Close() {
}
