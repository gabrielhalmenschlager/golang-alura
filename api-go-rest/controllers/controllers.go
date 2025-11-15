package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gabrielhalmenschlager/curso-golang-alura/api-go-rest/database"
	"github.com/gabrielhalmenschlager/curso-golang-alura/api-go-rest/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var personalidade []models.Personalidade
	database.DB.Find(&personalidade)
	json.NewEncoder(w).Encode(personalidade)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade []models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}
