import React, { useEffect, useState } from "react";
import PokemonCard from "./PokemonCard";

function Favoritos() {
  const [favoritos, setFavoritos] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8000/favoritos")
      .then((response) => response.json())
      .then((data) => setFavoritos(data))
      .catch(error => console.error("Erro ao buscar favoritos:", error));
  }, []);

  const removerFavorito = (id) => {
    fetch(`http://localhost:8000/favoritos/${id}`, {
      method: "DELETE",
    }).then(() => {
      setFavoritos(favoritos.filter((fav) => fav.pokemon.id !== id));
    }).catch(error => {
        console.error("Erro ao remover favorito:", error);
        alert("Falha ao remover Pokémon dos favoritos.");
    });
  };

  return (
    <div>
      <h2 className="page-title">Meus Favoritos</h2>
      <div className="list-grid">
        {favoritos.length === 0 && <p className="empty-list-text">Você ainda não tem Pokémon favoritos.</p>}
        
        {favoritos.map((fav) => (
          <PokemonCard
            key={fav.id}
            id={fav.pokemon.id}
            name={fav.pokemon.name}
            type={fav.pokemon.type}
            image_url={fav.pokemon.image_url}
            
            onRemove={() => removerFavorito(fav.pokemon.id)}
            isFavorited
          />
        ))}
      </div>
    </div>
  );
}

export default Favoritos;