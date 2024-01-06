package services

import (
	"database/sql"
	"log"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestDevolveScore(t *testing.T) {
	mockDB := &MockMySQLDatabase{
		MockConnectionString: "mock_connection_string",
		MockPonteiro:         &sql.DB{},
	}

	_, err := mockDB.InitDB()
	if err != nil {
		log.Fatal("Erro ao iniciar DB", err)
	}

	result := DevolveScore(mockDB)
	expectedResult := "Score 670 do user Primeiro\n"

	assert.Equal(t, expectedResult, result)
}
