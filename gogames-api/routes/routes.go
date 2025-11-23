package routes

import (
	"github.com/gabrielhalmenschlager/curso-golang-alura/gogames-api/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	config := cors.DefaultConfig()

	config.AllowOrigins = []string{"http://localhost:5173"}

	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}

	config.AllowHeaders = []string{"Origin", "Content-Type"}

	r.Use(cors.New(config))

	r.GET("/jogos", controllers.ListarJogos)
	r.GET("/jogos/:id", controllers.BuscarJogoPorID)
	r.POST("/jogos", controllers.CriarJogo)
	r.PUT("/jogos/:id", controllers.AtualizarJogo)
	r.DELETE("/jogos/:id", controllers.DeletarJogo)

	r.GET("/alugueis", controllers.ListarAlugueis)
	r.GET("/alugueis/ativos", controllers.ListarAlugueisAtivos)
	r.GET("/alugueis/inativos", controllers.ListarAlugueisInativos)
	r.POST("/alugueis/alugar/:id", controllers.AlugarJogo)
	r.PUT("/alugueis/devolver/:id", controllers.DevolverJogo)

	r.Run(":8000")
}
