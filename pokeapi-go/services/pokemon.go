package services

import (
	"errors"

	"github.com/gabrielhalmenschlager/curso-golang-alura/pokeapi-go/database"
	"github.com/gabrielhalmenschlager/curso-golang-alura/pokeapi-go/models"
)

func GetAllPokemons() ([]models.Pokemon, error) {
	var pokemons []models.Pokemon
	result := database.DB.Find(&pokemons)
	return pokemons, result.Error
}

func GetPokemonById(id string) (models.Pokemon, error) {
	var pokemon models.Pokemon
	result := database.DB.First(&pokemon, id)
	if result.Error != nil {
		return pokemon, errors.New("pokemon não encontrado")
	}
	return pokemon, nil
}

func CreatePokemon(pokemon models.Pokemon) (models.Pokemon, error) {
	result := database.DB.Create(&pokemon)
	return pokemon, result.Error
}
