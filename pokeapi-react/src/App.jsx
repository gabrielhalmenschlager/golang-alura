import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import PokemonList from "./components/PokemonList.jsx";
import Favoritos from "./components/Favoritos.jsx";
import AddPokemon from "./components/AddPokemon.jsx";
import EditPokemon from "./components/EditPokemon.jsx";
import Navbar from "./components/Navbar.jsx";
import "./App.css";

function App() {
  return (
    <BrowserRouter>
      <div className="pokedex-container">
        <Navbar />

        <main className="pokedex-screen">
          <Routes>
            <Route path="/" element={<PokemonList />} />
            <Route path="/favoritos" element={<Favoritos />} />
            <Route path="/adicionar" element={<AddPokemon />} />
            <Route path="/editar/:id" element={<EditPokemon />} />
          </Routes>
        </main>
      </div>
    </BrowserRouter>
  );
}

export default App;