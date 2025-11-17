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
	pokemonID := parsePokemonID(w, r)
	if pokemonID == 0 {
		return
	}

	favorito, err := services.CreateFavorito(pokemonID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(favorito)
}

func RemoveFavorito(w http.ResponseWriter, r *http.Request) {
	pokemonID := parsePokemonID(w, r)
	if pokemonID == 0 {
		return
	}

	if err := services.RemoveFavorito(pokemonID); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func parsePokemonID(w http.ResponseWriter, r *http.Request) uint {
	vars := mux.Vars(r)
	pokemonIDStr := vars["pokemonId"]

	var pokemonID uint
	_, err := fmt.Sscanf(pokemonIDStr, "%d", &pokemonID)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return 0
	}

	return pokemonID
}
