package main

import (
	"fmt"

	"github.com/gabrielhalmenschlager/curso-golang-alura/gogames-api/database"
	"github.com/gabrielhalmenschlager/curso-golang-alura/gogames-api/models"
	"github.com/gabrielhalmenschlager/curso-golang-alura/gogames-api/routes"
)

func main() {
	database.Conectar()

	database.DB.AutoMigrate(&models.Jogo{}, &models.Aluguel{})

	fmt.Println("Servidor rodando na porta 8000!")
	routes.HandleRequests()
}
