package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gabrielhalmenschlager/curso-golang-alura/pokeapi-go/models"
	"github.com/gabrielhalmenschlager/curso-golang-alura/pokeapi-go/services"
	"github.com/gorilla/mux"
)

func GetPokemons(w http.ResponseWriter, r *http.Request) {
	pokemons, err := services.GetAllPokemons()
	if err != nil {
		http.Error(w, "Erro ao listar pokemons", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(pokemons)
}

func GetPokemon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	pokemon, err := services.GetPokemonById(id)
	if err != nil {
		http.Error(w, "Pokemon não encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(pokemon)
}

func CreatePokemon(w http.ResponseWriter, r *http.Request) {
	var pokemon models.Pokemon
	if err := json.NewDecoder(r.Body).Decode(&pokemon); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	pokemon, err := services.CreatePokemon(pokemon)
	if err != nil {
		http.Error(w, "Erro ao salvar pokemon", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pokemon)
}
