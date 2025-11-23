package models

import "time"

type StatusAluguel string

const (
	ALUGADO   StatusAluguel = "ALUGADO"
	DEVOLVIDO StatusAluguel = "DEVOLVIDO"
)

type Aluguel struct {
	ID uint `gorm:"primaryKey" json:"id"`

	JogoID uint `json:"jogo_id"`
	Jogo   Jogo `json:"jogo"`

	DataAluguel           time.Time     `json:"data_aluguel" gorm:"autoCreateTime"`
	DataDevolucaoPrevista time.Time     `json:"data_devolucao_prevista"`
	DataDevolucaoReal     *time.Time    `json:"data_devolucao_real"`
	Status                StatusAluguel `json:"status" gorm:"default:ALUGADO"`
}
