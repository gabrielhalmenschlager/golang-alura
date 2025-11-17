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

func UpdatePokemon(id string, updatedData models.Pokemon) (models.Pokemon, error) {
	var pokemon models.Pokemon
	if err := database.DB.First(&pokemon, id).Error; err != nil {
		return pokemon, errors.New("pokemon não encontrado")
	}

	pokemon.Name = updatedData.Name
	pokemon.Type = updatedData.Type
	pokemon.ImageURL = updatedData.ImageURL

	if err := database.DB.Save(&pokemon).Error; err != nil {
		return pokemon, errors.New("erro ao atualizar pokemon")
	}

	return pokemon, nil
}

func DeletePokemon(id string) error {
	result := database.DB.Delete(&models.Pokemon{}, id)
	if result.RowsAffected == 0 {
		return errors.New("pokemon não encontrado")
	}
	return result.Error
}
