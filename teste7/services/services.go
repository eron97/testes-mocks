package services

import (
	"database/sql"
	"fmt"

	"github.com/eron97/testes-mocks.git/teste7/database"
	"github.com/eron97/testes-mocks.git/teste7/models"
)

type MockMySQLDatabase struct {
	MockConnectionString string
	MockPonteiro         *sql.DB
}

func (db *MockMySQLDatabase) InitDB() (*sql.DB, error) {
	return db.MockPonteiro, nil
}

func (db *MockMySQLDatabase) GetAllUsers() ([]models.User, error) {
	users := []models.User{
		{FirstName: "Primeiro"},
	}
	return users, nil
}

func DevolveScore(db database.Database) string {
	var result string

	users, _ := db.GetAllUsers()
	// Processa score dos usuários

	// Devolve score dos usuários

	for _, user := range users {
		result += fmt.Sprintf("Score 670 do user %s\n", user.FirstName)
	}

	return result
}
