package main

import "fmt"

type MockDatabase struct{}

func (db *MockDatabase) InitDB() error {
	return nil
}

func (db *MockDatabase) GetAllUsers() ([]string, error) {
	users := []string{"user1", "user2", "user3"}
	return users, nil
}

func main() {

	RealDatabase := &RealDatabase{}
	realProcess, realError := ProcessData(RealDatabase)
	fmt.Println("Dados reais do banco:", realProcess)
	fmt.Println("Erros reais de conexão:", realError)

	MockDatabase := &MockDatabase{}
	mockProcess, mockError := ProcessData(MockDatabase)
	fmt.Println("Dados mockados do banco:", mockProcess)
	fmt.Println("Erros mockados de conexão:", mockError)

}

type Database interface {
	InitDB() error
	GetAllUsers() ([]string, error)
}

type RealDatabase struct{}

func (db *RealDatabase) InitDB() error {
	return nil
}

func (db *RealDatabase) GetAllUsers() ([]string, error) {
	users := []string{"user1", "user2", "user3"}
	return users, nil
}

func ProcessData(db Database) ([]string, error) {
	if err := db.InitDB(); err != nil {
		return nil, err
	}

	users, err := db.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

/* OUTPUT TERMINAL

Dados reais do banco: [user1 user2 user3]
Erros reais de conexão: <nil>
Dados mockados do banco: [user1 user2 user3]
Erros mockados de conexão: <nil>

*/
