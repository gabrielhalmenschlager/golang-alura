package main

import (
	"fmt"

	"github.com/gabrielhalmenschlager/curso-golang-alura/estoque-go/internal/models"
)

func main() {
	fmt.Println("Sistema de Estoque")

	item1 := models.Item{
		ID:       1,
		Name:     "Primeiro Produto",
		Quantity: 10,
		Price:    19.99,
	}
	item2 := models.Item{
		ID:       2,
		Name:     "Segundo Produto",
		Quantity: 5,
		Price:    12.99,
	}
	item3 := models.Item{
		ID:       3,
		Name:     "Terceiro Produto",
		Quantity: 15,
		Price:    15.99,
	}

	fmt.Println(item1.Info())
	fmt.Println(item2.Info())
	fmt.Println(item3.Info())
}
