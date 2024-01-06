package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Database interface {
	InitDB() (*sql.DB, error)
	GetAllUsers(db *sql.DB) ([]User, error)
}

type MySQLDatabase struct {
	ConnectionString string
}

func main() {

	mysqlDB := &MySQLDatabase{
		ConnectionString: "admin:adminadmin@tcp(database-1.cpj0eavfzshu.us-east-1.rds.amazonaws.com:3306)/users",
	}

	connection, err := mysqlDB.InitDB()
	if err != nil {
		log.Fatal("Erro ao inicializar o banco de dados")
	}

	users, err := mysqlDB.GetAllUsers(connection)
	if err != nil {
		log.Fatal("Erro ao obter usuários:", err)
	}

	fmt.Println("Usuários do banco de dados:", users)

	defer connection.Close()

	MockmysqlDB := &MockMYSQLDatabase{
		MockConnectionstring: "MockConnection",
		MockUsers: []User{
			{ID: 1,
				Username:    "novo_usuario",
				Password:    "senha123",
				FirstName:   "Primeiro",
				LastName:    "Último",
				BirthDate:   "2000-01-01",
				PhoneNumber: "123456789",
			},
		},
	}

	mockconnection, err := MockmysqlDB.InitDB()
	mockusers, err := MockmysqlDB.GetAllUsers(mockconnection)
	if err != nil {
		log.Fatal("Erro ao obter users mocks:", err)
	}

	fmt.Println("Usuários do banco mock:", mockusers)

}

func (db *MySQLDatabase) InitDB() (*sql.DB, error) {
	// Criar conexão com o banco de dados MySQL
	connection, err := sql.Open("mysql", db.ConnectionString)
	if err != nil {
		return nil, err
	}

	// Verificar a conexão
	err = connection.Ping()
	if err != nil {
		return nil, err
	}

	return connection, nil
}

func (db *MySQLDatabase) GetAllUsers(connection *sql.DB) ([]User, error) {
	// Executar a consulta SQL para obter todos os usuários
	rows, err := connection.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	// Iterar sobre as linhas resultantes
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.BirthDate, &user.PhoneNumber)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		users = append(users, user)
	}

	// Verificar se ocorreram erros durante o processamento das linhas
	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return users, nil
}

type User struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	BirthDate   string `json:"birth_date"`
	PhoneNumber string `json:"phone_number"`
}

///////////////////// MOCKS

type MockMYSQLDatabase struct {
	MockConnectionstring string
	MockUsers            []User // Adicionando dados simulados
}

func (db *MockMYSQLDatabase) InitDB() (*sql.DB, error) {
	return nil, nil
}

func (db *MockMYSQLDatabase) GetAllUsers(connection *sql.DB) ([]User, error) {
	return db.MockUsers, nil
}

/* OUTPUT TERMINAL

Usuários do banco de dados: [{1 novo_usuario senha123 Primeiro Último 2000-01-01 123456789}]
Usuários do banco mock: [{1 novo_usuario senha123 Primeiro Último 2000-01-01 123456789}]

*/

// OBS: A partir do teste6 vamos levar mocks para serem usados dentro de funções com injeção de dependência onde realmente pode acontecer processamento interno dentro da função
