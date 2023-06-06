package main

import (
	"fmt"
	"github.com/mazzama/todo-grpc/cmd/server"
	"github.com/mazzama/todo-grpc/pkg/config"
	"github.com/mazzama/todo-grpc/pkg/database"
	"log"
)

func main() {
	err := config.NewAppConfig("./config", "config")
	if err != nil {
		panic(fmt.Errorf("failed to get app config: %v", err))
	}

	db, err := database.InitPostgres()
	if err != nil {
		log.Fatalln(err)
		return
	}

	serverReady := make(chan bool)
	s := server.Server{
		DBConn:      db,
		Port:        config.GetConfigInt("app.port"),
		ServerReady: serverReady,
	}

	s.Start()
}
