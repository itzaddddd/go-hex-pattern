package adapter

import (
	"errors"
	"testing"

	"github.com/itzaddddd/go-hex/core"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestSaveOrder(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error: %s", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	repo := NewOrderRepository(gormDB)

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "orders"`).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		err := repo.Save(core.Order{Total: 100.5})

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failure", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "orders"`).
			WillReturnError(errors.New("database error"))
		mock.ExpectRollback()

		err = repo.Save(core.Order{Total: 100})

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

}
