package database_test

import (
	"noahsTesting/backend/src/database"
	"testing"
)

func TestDataBase(t *testing.T) {
	db := database.New()

	db.Login("Noah")
	db.Login("Daan")
}
