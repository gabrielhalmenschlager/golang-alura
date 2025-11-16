package main

import (
	"github.com/gabrielhalmenschlager/curso-golang-alura/gin-api-rest/models"
	"github.com/gabrielhalmenschlager/curso-golang-alura/gin-api-rest/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Ana Silva", CPF: "123.456.789-00", RG: "12.345.678-9"},
		{Nome: "Bruno Santos", CPF: "987.654.321-00", RG: "98.765.432-1"},
		{Nome: "Carla Souza", CPF: "111.222.333-44", RG: "22.333.444-5"},
		{Nome: "Diego Martins", CPF: "555.666.777-88", RG: "55.666.777-8"},
		{Nome: "Eduarda Lima", CPF: "999.888.777-66", RG: "99.888.777-6"},
	}

	routes.HandleRequests()
}
