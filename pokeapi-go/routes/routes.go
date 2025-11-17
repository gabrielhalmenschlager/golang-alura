package routes

import (
	"log"
	"net/http"

	"github.com/gabrielhalmenschlager/curso-golang-alura/pokeapi-go/controllers"
	"github.com/gabrielhalmenschlager/curso-golang-alura/pokeapi-go/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequests() {
	r := mux.NewRouter()

	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/pokemon", controllers.GetPokemons).Methods("GET")
	r.HandleFunc("/pokemon/{id}", controllers.GetPokemon).Methods("GET")
	r.HandleFunc("/pokemon", controllers.CreatePokemon).Methods("POST")
	r.HandleFunc("/pokemon/{id}", controllers.UpdatePokemon).Methods("PUT")
	r.HandleFunc("/pokemon/{id}", controllers.DeletePokemon).Methods("DELETE")

	r.HandleFunc("/favoritos", controllers.GetFavoritos).Methods("GET")
	r.HandleFunc("/favoritos/{pokemonId}", controllers.CreateFavorito).Methods("POST")
	r.HandleFunc("/favoritos/{pokemonId}", controllers.RemoveFavorito).Methods("DELETE")
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(originsOk, headersOk, methodsOk)(r)))

}
