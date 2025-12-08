package routes

import (
	"myapi/internal/handlers"

	"github.com/gorilla/mux"
)

func CategoriaRoutes(r *mux.Router) {
	r.HandleFunc("/api/categorias", handlers.ListCategoriasHandler).Methods("GET")
	r.HandleFunc("/api/categorias/{id}", handlers.GetCategoriaHandler).Methods("GET")
	r.HandleFunc("/api/categorias", handlers.CreateCategoriaHandler).Methods("POST")
	r.HandleFunc("/api/categorias", handlers.UpdateCategoriaHandler).Methods("PUT")
	r.HandleFunc("/api/categorias/{id}", handlers.DeleteCategoriaHandler).Methods("DELETE")
}
