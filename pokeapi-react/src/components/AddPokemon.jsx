import { useState } from "react";

async function createPokemon(pokemon) {
  try {
    const response = await fetch("http://localhost:8000/pokemon", {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(pokemon),
    });
    if (!response.ok) {
      throw new Error('Falha ao criar Pokémon');
    }
    return await response.json();
  } catch (error) {
    console.error("Erro:", error);
  }
}


function AddPokemon() {
  const [pokemon, setPokemon] = useState({ name: "", type: "", image_url: "" });

  const handleChange = (e) => {
    setPokemon({ ...pokemon, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (!pokemon.name || !pokemon.type || !pokemon.image_url) {
        alert("Por favor, preencha todos os campos.");
        return;
    }
    await createPokemon(pokemon);
    alert("Pokémon adicionado com sucesso!");
    setPokemon({ name: "", type: "", image_url: "" });
  };

  return (
    <div className="form-container">
      <form onSubmit={handleSubmit} className="pokedex-form">
        <h2 className="page-title">Adicionar Novo Pokémon</h2>
        <div className="form-group">
          <label htmlFor="name">Nome:</label>
          <input id="name" name="name" placeholder="Ex: Pikachu" value={pokemon.name} onChange={handleChange} />
        </div>
        <div className="form-group">
          <label htmlFor="type">Tipo:</label>
          <input id="type" name="type" placeholder="Ex: electric" value={pokemon.type} onChange={handleChange} />
        </div>
        <div className="form-group">
          <label htmlFor="image_url">URL da Imagem:</label>
          <input id="image_url" name="image_url" placeholder="http://..." value={pokemon.image_url} onChange={handleChange} />
        </div>
        <button type="submit" className="btn submit-btn">Adicionar</button>
      </form>
    </div>
  );
}

export default AddPokemon;