package main

import (
	"gensmart/config"
	"gensmart/internal/infrastructure"
)

func main() {
	db := infrastructure.GetDBConnection()
	server := infrastructure.InitializeServer(db)

	server.Listen("127.0.0.1:" + config.AppConfig.Port)
}
