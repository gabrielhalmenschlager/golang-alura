package models

type Favorito struct {
	ID        uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	PokemonID uint    `json:"pokemon_id" gorm:"not null"`
	Pokemon   Pokemon `json:"pokemon" gorm:"foreignKey:PokemonID"`
}
