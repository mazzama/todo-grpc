package repository

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mazzama/todo-grpc/internal/todo/constant"
	"github.com/mazzama/todo-grpc/internal/todo/entity"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
	"time"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository ItemRepository
	person     *entity.Item
}

func (s *Suite) SetupSuite() {
	var db *sql.DB
	var err error

	db, s.mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(s.T(), err)

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})

	s.DB, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		s.Error(err, "failed to open Gorm DB")
	}

	if s.DB == nil {
		s.Error(err, "gorm DB is nil")
	}

	s.repository = NewItemRepository(s.DB)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_ItemRepository_Create() {
	var row = entity.Item{
		Name:        "Discrete Math",
		Description: "Travelling salesman problem",
		Notes:       "For next week",
		Status:      constant.ItemStatusTodo,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectQuery(`INSERT INTO "items" ("name","description","notes","status") VALUES ($1,$2,$3,$4) RETURNING "created_at","updated_at","id"`).
		WithArgs(row.Name, row.Description, row.Notes, row.Status).
		WillReturnRows(
			sqlmock.NewRows([]string{"created_at", "updated_at", "id"}).AddRow(time.Now(), time.Now(), 1))
	s.mock.ExpectCommit()

	err := s.repository.Create(context.Background(), &row)

	require.NoError(s.T(), err)
}

func (s *Suite) Test_ItemRepository_Update() {
	var row = entity.Item{
		ID:          1,
		Name:        "Discrete Math",
		Description: "Travelling salesman problem",
		Notes:       "For next week",
		Status:      constant.ItemStatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(`UPDATE "items" SET "name"=$1,"description"=$2,"notes"=$3,"status"=$4,"created_at"=$5,"updated_at"=$6 WHERE "id" = $7`).
		WithArgs(row.Name, row.Description, row.Notes, row.Status, AnyTime{}, AnyTime{}, row.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	_, err := s.repository.Update(context.Background(), &row)

	require.NoError(s.T(), err)
}

func (s *Suite) TestItemRepository_FindOneByCriteria() {
	s.mock.ExpectQuery(`SELECT * FROM "items" WHERE "id" = $1 ORDER BY "items"."id" LIMIT 1`).
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows([]string{"created_at", "updated_at", "id", "name", "description", "notes", "status"}).AddRow(time.Now(), time.Now(), 1, "Math", "Desc", "For tomorrow", "TODO"))

	_, err := s.repository.FindOneByCriteria(context.Background(), map[string]interface{}{
		"id": 1,
	})

	require.NoError(s.T(), err)
}

func (s *Suite) TestItemRepository_FindManyByCriteria() {
	s.mock.ExpectQuery(`SELECT * FROM "items" WHERE "status" = $1`).
		WithArgs(constant.ItemStatusTodo).
		WillReturnRows(
			sqlmock.NewRows([]string{"created_at", "updated_at", "id", "name", "description", "notes", "status"}).AddRow(time.Now(), time.Now(), 1, "Math", "Desc", "For tomorrow", "TODO"))

	_, err := s.repository.FindManyByCriteria(context.Background(), map[string]interface{}{
		"status": constant.ItemStatusTodo,
	})

	require.NoError(s.T(), err)
}

func (s *Suite) TestItemRepository_Delete() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(`DELETE FROM "items" WHERE id = $1`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.repository.Delete(context.Background(), 1)

	require.NoError(s.T(), err)
}
