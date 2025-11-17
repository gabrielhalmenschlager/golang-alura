package routes

import (
	"net/http"

	"github.com/gabrielhalmenschlager/curso-golang-alura/pokeapi-go/controllers"
	"github.com/gabrielhalmenschlager/curso-golang-alura/pokeapi-go/middleware"
	"github.com/gorilla/mux"
)

func HandleRequests() {
	r := mux.NewRouter()

	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/pokemon", controllers.GetPokemons).Methods("GET")
	r.HandleFunc("/pokemon/{id}", controllers.GetPokemon).Methods("GET")
	r.HandleFunc("/pokemon", controllers.CreatePokemon).Methods("POST")

	r.HandleFunc("/favoritos", controllers.GetFavoritos).Methods("GET")
	r.HandleFunc("/favoritos/{pokemonId}", controllers.CreateFavorito).Methods("POST")

	http.ListenAndServe(":8000", r)
}
