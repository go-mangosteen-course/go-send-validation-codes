package cmd

import (
	"log"
	"mangosteen/internal/database"
	"mangosteen/internal/router"
)

func RunServer() {
	database.MysqlConnect()
	database.MysqlCreateTable()
	defer database.MysqlClose()
	r := router.New()
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("r.Run 的下一行")

}
