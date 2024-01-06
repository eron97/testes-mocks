package services

import (
	"database/sql"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
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
