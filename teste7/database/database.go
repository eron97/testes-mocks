package database

import (
	"database/sql"
	"log"

	"github.com/eron97/testes-mocks.git/teste7/models"
	_ "github.com/go-sql-driver/mysql"
)

type Database interface {
	InitDB() (*sql.DB, error)
	GetAllUsers() ([]models.User, error)
}

type MySQLDatabase struct {
	ConnectionString string
	Ponteiro         *sql.DB
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

	db.Ponteiro = connection // Atribuir a conexão ao campo Ponteiro

	return connection, nil
}

func (db *MySQLDatabase) GetAllUsers() ([]models.User, error) {
	// Executar a consulta SQL para obter todos os usuários
	rows, err := db.Ponteiro.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	// Iterar sobre as linhas resultantes
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.BirthDate, &user.PhoneNumber)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
