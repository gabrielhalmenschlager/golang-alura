package controllers

import (
	"net/http"
	"strconv"

	"github.com/gabrielhalmenschlager/curso-golang-alura/gogames-api/services"
	"github.com/gin-gonic/gin"
)

type AlugarRequest struct {
	JogoID uint `json:"jogo_id"`
}

func AlugarJogo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	aluguel, erro := services.AlugarJogo(uint(id))
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": erro.Error()})
		return
	}

	c.JSON(http.StatusCreated, aluguel)
}

func DevolverJogo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	aluguel, err := services.DevolverJogo(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, aluguel)
}

func ListarAlugueis(c *gin.Context) {
	alugueis, err := services.BuscarAlugueis("TODOS")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}
	c.JSON(http.StatusOK, alugueis)
}

func ListarAlugueisAtivos(c *gin.Context) {
	alugueis, err := services.BuscarAlugueis("ALUGADO")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}
	c.JSON(http.StatusOK, alugueis)
}

func ListarAlugueisInativos(c *gin.Context) {
	alugueis, err := services.BuscarAlugueis("DEVOLVIDO")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}
	c.JSON(http.StatusOK, alugueis)
}
