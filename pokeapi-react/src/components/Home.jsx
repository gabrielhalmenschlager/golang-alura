import { useEffect, useState } from "react";
import PokemonCard from "./PokemonCard";
import { getPokemons } from "../api";

function Home() {
  const [pokemons, setPokemons] = useState([]);

  useEffect(() => {
    getPokemons().then(data => setPokemons(data));
  }, []);

  return (
    <div>
      <h1>Pokédex</h1>
      <div className="pokemon-container">
        {pokemons.map((pokemon) => (
          <PokemonCard key={pokemon.id} {...pokemon} />
        ))}
      </div>
    </div>
  );
}

export default Home;
