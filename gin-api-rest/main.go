package main

import (
	"github.com/gabrielhalmenschlager/curso-golang-alura/gin-api-rest/database"
	"github.com/gabrielhalmenschlager/curso-golang-alura/gin-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
