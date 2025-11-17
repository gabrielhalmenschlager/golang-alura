const API_BASE_URL = "http://localhost:8000";

export const getPokemons = async () => {
  const response = await fetch(`${API_BASE_URL}/pokemon`);
  return response.json();
};

export const createPokemon = async (pokemon) => {
  const response = await fetch(`${API_BASE_URL}/pokemon`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(pokemon),
  });
  return response.json();
};
