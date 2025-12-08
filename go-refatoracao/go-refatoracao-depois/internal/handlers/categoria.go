package handlers

import (
	"encoding/json"
	"myapi/internal/config"
	"myapi/internal/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ListCategoriasHandler(w http.ResponseWriter, r *http.Request) {
	var categorias []models.Categoria
	if err := config.DB.Find(&categorias).Error; err != nil {
		http.Error(w, "Erro ao buscar categorias", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(categorias)
}

func GetCategoriaHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "ID não fornecido", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	var categoria models.Categoria
	if err := config.DB.First(&categoria, id).Error; err != nil {
		http.Error(w, "Categoria não encontrada", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(categoria)
}

func CreateCategoriaHandler(w http.ResponseWriter, r *http.Request) {
	var categoria models.Categoria
	if err := json.NewDecoder(r.Body).Decode(&categoria); err != nil {
		http.Error(w, "Erro ao decodificar a categoria", http.StatusBadRequest)
		return
	}
	if err := config.DB.Create(&categoria).Error; err != nil {
		http.Error(w, "Erro ao criar a categoria", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(categoria)
}

func UpdateCategoriaHandler(w http.ResponseWriter, r *http.Request) {
	var categoria models.Categoria
	if err := json.NewDecoder(r.Body).Decode(&categoria); err != nil {
		http.Error(w, "Erro ao decodificar a categoria", http.StatusBadRequest)
		return
	}
	if err := config.DB.Save(&categoria).Error; err != nil {
		http.Error(w, "Erro ao atualizar a categoria", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(categoria)
}

func DeleteCategoriaHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "ID não fornecido", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	if err := config.DB.Delete(&models.Categoria{}, id).Error; err != nil {
		http.Error(w, "Erro ao deletar a categoria", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Categoria deletada com sucesso"))
}
