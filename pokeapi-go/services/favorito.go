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
	var existingFav models.Favorito

	if err := database.DB.Where("pokemon_id = ?", pokemonID).First(&existingFav).Error; err == nil {
		return existingFav, errors.New("Pokémon já está nos favoritos")
	}

	favorito := models.Favorito{PokemonID: pokemonID}
	result := database.DB.Create(&favorito)
	if result.Error != nil {
		return favorito, errors.New("erro ao salvar favorito")
	}
	return favorito, nil
}

func RemoveFavorito(pokemonID uint) error {
	var favorito models.Favorito
	if err := database.DB.Where("pokemon_id = ?", pokemonID).First(&favorito).Error; err != nil {
		return errors.New("favorito não encontrado")
	}

	return database.DB.Delete(&favorito).Error
}
