package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gabrielhalmenschlager/curso-golang-alura/pokeapi-go/services"
	"github.com/gorilla/mux"
)

func GetFavoritos(w http.ResponseWriter, r *http.Request) {
	favoritos, err := services.GetAllFavoritos()
	if err != nil {
		http.Error(w, "Erro ao listar favoritos", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(favoritos)
}

func CreateFavorito(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pokemonIDStr := vars["pokemonId"]

	var pokemonID uint
	fmt.Sscanf(pokemonIDStr, "%d", &pokemonID)

	favorito, err := services.CreateFavorito(pokemonID)
	if err != nil {
		http.Error(w, "Erro ao favoritar pokemon", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(favorito)
}
