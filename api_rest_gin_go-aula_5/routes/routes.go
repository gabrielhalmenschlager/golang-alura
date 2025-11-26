package routes

import (
	"github.com/gabrielhalmenschlager/curso-golang-alura/api_rest_gin_go-aula_5/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("api_rest_gin_go-aula_5/templates/*")
	r.Static("api_rest_gin_go-aula_5/assets", "./api_rest_gin_go-aula_5/assets")
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.GET("/index", controllers.ExibePaginaIndex)
	r.NoRoute(controllers.RotaNaoEncontrada)
	r.Run()
}
