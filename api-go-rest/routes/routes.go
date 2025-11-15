package routes

import (
	"log"
	"net/http"

	"github.com/gabrielhalmenschlager/curso-golang-alura/api-go-rest/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/personalidades", controllers.TodasPersonalidades)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
