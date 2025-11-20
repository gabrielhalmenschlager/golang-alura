package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gabrielhalmenschlager/curso-golang-alura/api_rest_gin_go-aula_5/controllers"
	"github.com/gin-gonic/gin"
)

func SetupDasRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

// func TestFalhador(t *testing.T) {
// 	t.Fatalf("Teste falhou de propósito, não se preocupe!")
// }

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/gabe", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	if resposta.Code != http.StatusOK {
		t.Fatalf("Status error: valor recebido foi %d e o esperado era %d", resposta.Code, http.StatusOK)
	}
}
