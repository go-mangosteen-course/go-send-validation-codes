package cmd

import (
	"log"
	"mangosteen/internal/database"
	"mangosteen/internal/router"
)

func RunServer() {
	database.Connect()
	database.CreateTables()
	defer database.Close()
	r := router.New()
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("r.Run 的下一行")

}
