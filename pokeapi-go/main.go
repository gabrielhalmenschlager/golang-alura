package main

import (
	"fmt"

	"github.com/gabrielhalmenschlager/curso-golang-alura/pokeapi-go/database"
	"github.com/gabrielhalmenschlager/curso-golang-alura/pokeapi-go/models"
	"github.com/gabrielhalmenschlager/curso-golang-alura/pokeapi-go/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	database.DB.AutoMigrate(&models.Pokemon{}, &models.Favorito{})

	fmt.Println("Servidor rodando na porta 8000")
	routes.HandleRequests()
}
