package main

import (
	"context"
	"fmt"
	"github.com/mazzama/todo-grpc/internal/todo/handler"
	"github.com/mazzama/todo-grpc/internal/todo/repository"
	"github.com/mazzama/todo-grpc/internal/todo/service"
	"github.com/mazzama/todo-grpc/pkg/config"
	"github.com/mazzama/todo-grpc/pkg/database"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
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

	itemRepository := repository.NewItemRepository(db)
	itemService := service.NewItemService(itemRepository)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalln(err)
		return
	}

	server := grpc.NewServer()
	handler.NewTodoServerGrpc(server, itemService)

	// Start the server in a separate goroutine
	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	log.Println("Server started")

	<-ctx.Done()
	server.GracefulStop()

	log.Println("Server stopped gracefully")
}
