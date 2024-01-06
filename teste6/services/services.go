package services

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/eron97/testes-mocks.git/teste6/database"
)

func DevolveScore(mysqlDB *database.MySQLDatabase, connection *sql.DB) {
	users, err := mysqlDB.GetAllUsers(connection)
	if err != nil {
		log.Fatal("Ponteiro ou dados não existem", err)
	}

	// Processa score dos usuários

	// Devolve score dos usuários

	for _, user := range users {
		fmt.Printf("Score 670 do user %s", user.FirstName)
	}

}
