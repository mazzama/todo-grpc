package server

import (
	"context"
	"fmt"
	"github.com/mazzama/todo-grpc/internal/todo/handler"
	"github.com/mazzama/todo-grpc/internal/todo/repository"
	"github.com/mazzama/todo-grpc/internal/todo/service"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"net"
	"os"
	"os/signal"
)

type Server struct {
	Port        int
	DBConn      *gorm.DB
	ServerReady chan bool
	Server      *grpc.Server
}

// Start start grpc server
func (s *Server) Start() {
	itemRepository := repository.NewItemRepository(s.DBConn)
	itemService := service.NewItemService(itemRepository)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", s.Port))
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

		if s.ServerReady != nil {
			s.Server = server
			s.ServerReady <- true
		}
	}()

	log.Println("Server started")

	<-ctx.Done()
	server.GracefulStop()

	log.Println("Server stopped gracefully")
}
