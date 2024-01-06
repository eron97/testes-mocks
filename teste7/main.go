package main

import (
	"fmt"
	"log"

	"github.com/eron97/testes-mocks.git/teste7/database"
	"github.com/eron97/testes-mocks.git/teste7/services"
)

func main() {

	mysqlDB := &database.MySQLDatabase{
		ConnectionString: "admin:adminadmin@tcp(database-1.cpj0eavfzshu.us-east-1.rds.amazonaws.com:3306)/users",
	}

	_, err := mysqlDB.InitDB()
	if err != nil {
		log.Fatal("Erro ao iniciar DB", err)
	}

	fmt.Println("Banco de dados ON")

	services.DevolveScore(mysqlDB)

	mysqlDB.Ponteiro.Close()

	fmt.Println("Banco de dados OFF")
}
