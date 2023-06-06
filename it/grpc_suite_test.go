package it

import (
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/mazzama/todo-grpc/cmd/server"
	"github.com/mazzama/todo-grpc/internal/todo/handler"
	"github.com/mazzama/todo-grpc/internal/todo/repository"
	"github.com/mazzama/todo-grpc/internal/todo/service"
	"github.com/mazzama/todo-grpc/pkg/config"
	"github.com/mazzama/todo-grpc/pkg/database"
	"github.com/mazzama/todo-grpc/pkg/pb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"gorm.io/gorm"
	"log"
	"net"
	"os"
	"syscall"
	"testing"
)

var lis *bufconn.Listener

type e2eTestSuite struct {
	suite.Suite
	dbConnectionStr string
	port            int
	dbConn          *gorm.DB
	dbMigration     *migrate.Migrate
}

func TestE2ETestSuite(t *testing.T) {
	suite.Run(t, &e2eTestSuite{})
}

func (s *e2eTestSuite) SetupSuite() {
	err := config.NewAppConfig("./../config", "config")
	if err != nil {
		panic(fmt.Errorf("failed to get app config: %v", err))
	}

	s.port = config.GetConfigInt("app.port")
	s.dbConnectionStr = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.GetConfigString(config.PostgresUser),
		config.GetConfigString(config.PostgresPassword),
		config.GetConfigString(config.PostgresHost),
		config.GetConfigString(config.PostgresPort),
		config.GetConfigString(config.PostgresDBName),
	)

	migration, err := migrate.New("file://../migration", s.dbConnectionStr)
	s.Require().NoError(err)

	s.dbMigration = migration
}

func (s *e2eTestSuite) TearDownSuite() {
	p, _ := os.FindProcess(syscall.Getpid())
	p.Signal(syscall.SIGINT)
}

func (s *e2eTestSuite) SetupTest() {
	log.Println("Start the migration")
	if err := s.dbMigration.Up(); err != nil && err != migrate.ErrNoChange {
		s.Require().NoError(err)
	}
}

func (s *e2eTestSuite) TearDownTest() {
	m, err := migrate.New("file://../migration", s.dbConnectionStr)
	log.Println("Tear down the migration")

	assert.NoError(s.T(), err)
	assert.NoError(s.T(), m.Down())
}

func bufDialer(context.Context, string) (net.Conn, error) {
	var ll = lis
	return ll.Dial()
}
func setupServer(ctx context.Context, port int) (pb.TodoServiceClient, func()) {
	buffer := 101024 * 1024
	lis = bufconn.Listen(buffer)

	db, err := database.InitPostgres()
	if err != nil {
		log.Fatalln(err)
		return nil, nil
	}

	serverReady := make(chan bool)
	s := server.Server{
		DBConn:      db,
		Port:        config.GetConfigInt("app.port"),
		ServerReady: serverReady,
	}

	itemRepository := repository.NewItemRepository(s.DBConn)
	itemService := service.NewItemService(itemRepository)

	appServer := grpc.NewServer()
	handler.NewTodoServerGrpc(appServer, itemService)
	go func() {
		if err := appServer.Serve(lis); err != nil {
			log.Printf("Server exited with error: %v", err)
		}
	}()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial bufnet: %v", err)
	}
	client := pb.NewTodoServiceClient(conn)

	closer := func() {
		log.Println("Closing the connection")
	}

	return client, closer
}
