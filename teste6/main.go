package main

import (
	"fmt"
	"log"

	"github.com/eron97/testes-mocks.git/teste6/database"
	"github.com/eron97/testes-mocks.git/teste6/services"
)

func main() {

	mysqlDB := &database.MySQLDatabase{
		ConnectionString: "admin:adminadmin@tcp(database-1.cpj0eavfzshu.us-east-1.rds.amazonaws.com:3306)/users",
	}

	connection, err := mysqlDB.InitDB()
	if err != nil {
		log.Fatal("Erro ao iniciar BD: ", err)
	}

	fmt.Println("Banco de dados ON")

	services.DevolveScore(mysqlDB, connection)

	defer connection.Close()

}

/* OUTPUT TERMINAL

Banco de dados ON
Score 670 do user Primeiro

*/
