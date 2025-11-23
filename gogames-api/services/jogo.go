package services

import (
	"errors"

	"github.com/gabrielhalmenschlager/curso-golang-alura/gogames-api/database"
	"github.com/gabrielhalmenschlager/curso-golang-alura/gogames-api/models"
)

func CriarJogo(jogo models.Jogo) (models.Jogo, error) {
	result := database.DB.Create(&jogo)
	return jogo, result.Error
}

func ListarJogos() ([]models.Jogo, error) {
	var jogos []models.Jogo
	result := database.DB.Find(&jogos)
	return jogos, result.Error
}

func BuscarJogoPorID(id uint) (models.Jogo, error) {
	var jogo models.Jogo
	result := database.DB.First(&jogo, id)

	if result.RowsAffected == 0 {
		return jogo, errors.New("jogo não encontrado")
	}

	return jogo, nil
}

func AtualizarJogo(id uint, dados models.Jogo) (models.Jogo, error) {
	jogo, err := BuscarJogoPorID(id)
	if err != nil {
		return jogo, err
	}

	jogo.Nome = dados.Nome
	jogo.Descricao = dados.Descricao
	jogo.Categoria = dados.Categoria
	jogo.ImagemURL = dados.ImagemURL

	result := database.DB.Save(&jogo)
	return jogo, result.Error
}

func DeletarJogo(id uint) error {
	jogo, err := BuscarJogoPorID(id)
	if err != nil {
		return err
	}

	if !jogo.Disponivel {
		return errors.New("não é possível excluir um jogo alugado")
	}

	result := database.DB.Delete(&jogo)
	return result.Error
}
