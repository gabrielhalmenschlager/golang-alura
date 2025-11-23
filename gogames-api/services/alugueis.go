package services

import (
	"errors"
	"time"

	"github.com/gabrielhalmenschlager/curso-golang-alura/gogames-api/database"
	"github.com/gabrielhalmenschlager/curso-golang-alura/gogames-api/models"
)

func calcularPrazo(categoria models.Categoria) int {
	switch categoria {
	case models.BRONZE:
		return 3
	case models.PRATA:
		return 5
	case models.OURO:
		return 7
	default:
		return 3
	}
}

func AlugarJogo(jogoID uint) (models.Aluguel, error) {
	var aluguel models.Aluguel

	jogo, err := BuscarJogoPorID(jogoID)
	if err != nil {
		return aluguel, err
	}

	if !jogo.Disponivel {
		return aluguel, errors.New("jogo não está disponível para aluguel")
	}

	dias := calcularPrazo(jogo.Categoria)

	aluguel = models.Aluguel{
		JogoID:                jogo.ID,
		DataAluguel:           time.Now(),
		DataDevolucaoPrevista: time.Now().AddDate(0, 0, dias),
		Status:                models.ALUGADO,
	}

	result := database.DB.Create(&aluguel)
	if result.Error != nil {
		return aluguel, result.Error
	}

	jogo.Disponivel = false
	database.DB.Save(&jogo)

	return aluguel, nil
}

func DevolverJogo(aluguelID uint) (models.Aluguel, error) {
	var aluguel models.Aluguel

	result := database.DB.Preload("Jogo").First(&aluguel, aluguelID)
	if result.RowsAffected == 0 {
		return aluguel, errors.New("aluguel não encontrado")
	}

	if aluguel.Status == models.DEVOLVIDO {
		return aluguel, errors.New("este aluguel já foi devolvido")
	}

	now := time.Now()
	aluguel.DataDevolucaoReal = &now
	aluguel.Status = models.DEVOLVIDO

	database.DB.Save(&aluguel)

	aluguel.Jogo.Disponivel = true
	database.DB.Save(&aluguel.Jogo)

	return aluguel, nil
}

func ListarAlugueis() ([]models.Aluguel, error) {
	var alugueis []models.Aluguel
	result := database.DB.Preload("Jogo").Find(&alugueis)
	return alugueis, result.Error
}

func BuscarAlugueis(status string) ([]models.Aluguel, error) {
	var alugueis []models.Aluguel

	query := database.DB.Preload("Jogo")

	if status != "TODOS" {
		query = query.Where("status = ?", status)
	}

	if err := query.Find(&alugueis).Error; err != nil {
		return nil, err
	}

	return alugueis, nil
}
