package routes

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

/*

Teste de Integração Automatizado

verifica o comportamento de duas unidades de código juntas
roteador Gin - funções manipuladoras de rotas

Cliente HTTP envia requisições para o roteador e verifica status code

*/

func TestValidarRotas(t *testing.T) {

	router := gin.Default()
	SetupRoutes(router)

	rotasValidas := map[string]string{
		"/":         "GET",
		"/register": "POST",
	}

	for rota, metodo := range rotasValidas {
		req, _ := http.NewRequest(metodo, rota, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			fmt.Printf("Rota %s: Status code inesperado: %d\n", rota, w.Code)
			t.Fail()
		}
	}
}

func TestWelcomeHandler(t *testing.T) {
	// Cria um novo contexto de teste
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	WelcomeHandler(c)

	if w.Code != http.StatusOK {
		t.Errorf("Status code inesperado: %d", w.Code)
	}

	body, _ := ioutil.ReadAll(w.Body)
	if string(body) != `{"message":"Bem-vindo ao servidor Gin!"}` {
		t.Errorf("Mensagem de resposta incorreta: %s", string(body))
	}
}
