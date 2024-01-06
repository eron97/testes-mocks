package api

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
