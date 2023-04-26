package main

import (
	"log"
	"mangosteen/cmd"

	"github.com/spf13/viper"
)

// @title           山竹记账 API
// @description     这是一个使用 Swagger 2.0 标准编写的 API 文档。

// @contact.name   方应杭
// @contact.url    https://fangyinghang.com
// @contact.email  fangyinghang@foxmail.com

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth(JWT)

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
	cmd.Run()
}
