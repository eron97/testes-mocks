package main

// criando mock APIClient
type MockAPIClient struct{}

func (c *MockAPIClient) Get(url string) (string, error) {
	return "Resposta do mock API", nil
}

func main() {
	// Usando implementação padrão

}
