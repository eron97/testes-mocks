package main

import "fmt"

// criando mock APIClient
type MockAPIClient struct{}

func (c *MockAPIClient) Get(url string) (string, error) {
	return "Resposta do mock API", nil
}

func main() {
	// Usando implementação padrão
	realClient := &DefaultAPIClient{}
	realData, realErr := GetDataFromAPI(realClient)
	fmt.Println("Dados reais:", realData)
	fmt.Println("Erros reais:", realErr)

	// Usando mock para teste
	mockClient := &MockAPIClient{}
	mockData, mockErr := GetDataFromAPI(mockClient)
	fmt.Println("Dados mocks:", mockData)
	fmt.Println("Erros mocks:", mockErr)

}

type APIClient interface {
	Get(url string) (string, error)
}

type DefaultAPIClient struct{}

func (c *DefaultAPIClient) Get(url string) (string, error) {

	return "Resposta da API real", nil
}

// função que utiliza o cliente da api para obter dados

func GetDataFromAPI(client APIClient) (string, error) {
	response, err := client.Get("https://api.exemplo.com/dados")
	if err != nil {
		return "", err
	}

	return response, nil
}
