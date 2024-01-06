package services

import (
	"fmt"

	"github.com/eron97/testes-mocks.git/teste8/database"
)

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
