import React, { useEffect, useState, useMemo } from "react";
import PokemonCard from "./PokemonCard";

function PokemonList() {
  const [pokemons, setPokemons] = useState([]);
  const [favoritos, setFavoritos] = useState([]);

  useEffect(() => {
    const fetchPokemons = fetch("http://localhost:8000/pokemon").then((res) => res.json());
    const fetchFavoritos = fetch("http://localhost:8000/favoritos").then((res) => res.json());

    Promise.all([fetchPokemons, fetchFavoritos])
      .then(([pokemonData, favoritoData]) => {
        setPokemons(pokemonData);
        setFavoritos(favoritoData);
      })
      .catch(error => console.error("Erro ao buscar dados:", error));
  }, []);

  const favoritoIds = useMemo(() =>
    new Set(favoritos.map((fav) => fav.pokemon.id))
  , [favoritos]);

  const adicionarFavorito = (id) => {
    fetch(`http://localhost:8000/favoritos/${id}`, {
      method: "POST",
    }).then(() => {
      const pokemonAdicionado = pokemons.find(p => p.id === id);
      if (pokemonAdicionado) {
        setFavoritos([...favoritos, { id: Math.random(), pokemon: pokemonAdicionado }]);
      }
    });
  };

  const removerFavorito = (id) => {
    fetch(`http://localhost:8000/favoritos/${id}`, {
      method: "DELETE",
    }).then(() => {
      setFavoritos(favoritos.filter((fav) => fav.pokemon.id !== id));
    });
  };

  const handleDelete = (id) => {
    if (window.confirm(`Tem certeza que deseja deletar o Pokémon #${id}?`)) {
        fetch(`http://localhost:8000/pokemon/${id}`, {
            method: "DELETE",
        }).then(() => {
            setPokemons(pokemons.filter((poke) => poke.id !== id));
            setFavoritos(favoritos.filter((fav) => fav.pokemon.id !== id));
        }).catch(error => console.error("Erro ao deletar Pokémon:", error));
    }
};

  return (
    <div>
      <h2 className="page-title">Pokémon Encontrados</h2>
      <div className="list-grid">
        {pokemons.map((poke) => {
          const isFavorited = favoritoIds.has(poke.id);

          return (
            <PokemonCard
              key={poke.id}
              id={poke.id}
              name={poke.name}
              type={poke.type}
              image_url={poke.image_url}
              
              isFavorited={isFavorited}
              onFavorite={adicionarFavorito}
              onRemove={removerFavorito}
              onDelete={handleDelete}
              showAdminActions={true}
            />
          );
        })}
      </div>
    </div>
  );
}

export default PokemonList;