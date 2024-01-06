package services

import (
	"fmt"
	"testing"

	mocks "github.com/eron97/testes-mocks.git/teste8/database/gomock"
	"github.com/eron97/testes-mocks.git/teste8/models"
	"go.uber.org/mock/gomock"
)

func TestDevolveScore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDatabase(ctrl)

	expectedUsers := []models.User{
		{ID: 1, Username: "novo_usuario", Password: "senha123", FirstName: "Primeiro", LastName: "Último", BirthDate: "2000-01-01", PhoneNumber: "123456789"},
	}
	mockDB.EXPECT().GetAllUsers().Return(expectedUsers, nil)

	result := DevolveScore(mockDB)

	expectedResult := "Score 670 do user Primeiro\n"

	if result != expectedResult {
		t.Errorf("Expected result: %s, got: %s", expectedResult, result)
	}

	fmt.Printf("result test mock: %s", result)
}
