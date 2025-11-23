package controllers

import (
	"net/http"
	"strconv"

	"github.com/gabrielhalmenschlager/curso-golang-alura/gogames-api/models"
	"github.com/gabrielhalmenschlager/curso-golang-alura/gogames-api/services"
	"github.com/gin-gonic/gin"
)

func CriarJogo(c *gin.Context) {
	var jogo models.Jogo

	if err := c.ShouldBindJSON(&jogo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	resultado, err := services.CriarJogo(jogo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resultado)
}

func ListarJogos(c *gin.Context) {
	jogos, err := services.ListarJogos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, jogos)
}

func BuscarJogoPorID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	jogo, err := services.BuscarJogoPorID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, jogo)
}

func AtualizarJogo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var dados models.Jogo
	if err := c.ShouldBindJSON(&dados); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	jogo, err := services.AtualizarJogo(uint(id), dados)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, jogo)
}

func DeletarJogo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := services.DeletarJogo(uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "jogo deletado com sucesso"})
}
