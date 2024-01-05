package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestWelcomeEndpoint(t *testing.T) {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Bem-vindo ao servidor Gin!",
		})
	})

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Esperado 200, obtido %d", w.Code)
	}

	expected := `{"message":"Bem-vindo ao servidor Gin!"}`
	if w.Body.String() != expected {
		t.Errorf("Resposta inesperada: %s", w.Body.String())
	}
}
