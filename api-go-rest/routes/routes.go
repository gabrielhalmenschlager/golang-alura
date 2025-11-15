package routes

import (
	"log"
	"net/http"

	"github.com/gabrielhalmenschlager/curso-golang-alura/api-go-rest/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalidades", controllers.TodasPersonalidades)
	log.Fatal(http.ListenAndServe(":8000", r))
}
