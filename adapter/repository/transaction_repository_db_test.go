package repository

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vitorath/clean-architecture-golang/adapter/repository/fixtures"
)

func TestTransactionRepositoryDbInsert(t *testing.T) {
	migrationsDir := os.DirFS("fixtures/sql")
	db := fixtures.Up(migrationsDir)
	defer fixtures.Down(db, migrationsDir)

	repository := NewTransactionRepositoryDb(db)
	err := repository.Insert("1", "1", 2, "approved", "")
	assert.Nil(t, err)
}
