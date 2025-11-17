package services

import (
	"errors"

	"github.com/gabrielhalmenschlager/curso-golang-alura/pokeapi-go/database"
	"github.com/gabrielhalmenschlager/curso-golang-alura/pokeapi-go/models"
)

func GetAllFavoritos() ([]models.Favorito, error) {
	var favoritos []models.Favorito
	result := database.DB.Preload("Pokemon").Find(&favoritos)
	return favoritos, result.Error
}

func CreateFavorito(pokemonID uint) (models.Favorito, error) {
	var favorito models.Favorito
	favorito.PokemonID = pokemonID

	result := database.DB.Create(&favorito)
	if result.Error != nil {
		return favorito, errors.New("erro ao salvar favorito")
	}
	return favorito, nil
}
