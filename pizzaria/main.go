package main

import (
	"github.com/gin-gonic/gin"
)

type Pizza struct {
	ID    int     `json:"id"`
	Nome  string  `json:"nome"`
	Preco float64 `json:"preco"`
}

func main() {
	router := gin.Default()
	router.GET("/pizzas", getPizzas)
	router.Run()
}

func getPizzas(c *gin.Context) {
	pizzas := []Pizza{
		{ID: 1, Nome: "Toscana", Preco: 79.5},
		{ID: 2, Nome: "Marguerita", Preco: 69.5},
		{ID: 3, Nome: "Atum com queijo", Preco: 59.5},
	}
	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
}
