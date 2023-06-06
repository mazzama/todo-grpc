package it

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/mazzama/todo-grpc/pkg/config"
	"github.com/mazzama/todo-grpc/pkg/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"log"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var (
	connStr string
)

type PostgresRepositoryTestSuite struct {
	gormDB *gorm.DB
	suite.Suite
}

func (p *PostgresRepositoryTestSuite) SetupSuite() {
	err := config.NewAppConfig("./../config", "config")
	if err != nil {
		panic(fmt.Errorf("failed to get app config: %v", err))
	}

	connStr = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.GetConfigString(config.PostgresUser),
		config.GetConfigString(config.PostgresPassword),
		config.GetConfigString(config.PostgresHost),
		config.GetConfigString(config.PostgresPort),
		config.GetConfigString(config.PostgresDBName),
	)

	fmt.Println(connStr)

	p.gormDB, err = database.InitPostgres()
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func TestPostgresRepositoryTestSuite(t *testing.T) {
	suite.Run(t, &PostgresRepositoryTestSuite{})
}

func (p *PostgresRepositoryTestSuite) SetupTest() {
	m, err := migrate.New("file://../migration", connStr)
	assert.NoError(p.T(), err)

	log.Println("Apply the migration")

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return
		}

		panic(err)
	}
}

func (p *PostgresRepositoryTestSuite) TearDownTest() {
	m, err := migrate.New("file://../migration", connStr)
	log.Println("Tear down the migration")

	assert.NoError(p.T(), err)
	assert.NoError(p.T(), m.Down())
}
