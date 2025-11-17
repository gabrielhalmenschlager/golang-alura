import React from "react";
import { NavLink } from "react-router-dom";
import "./Navbar.css";

function Navbar() {
  return (
    <nav className="pokedex-nav">
      <div className="pokedex-nav-lights">
        <div className="light big-light-blue"></div>
        <div className="light small-light-red"></div>
        <div className="light small-light-yellow"></div>
        <div className="light small-light-green"></div>
      </div>
      <div className="nav-links">
        <NavLink
          to="/"
          className={({ isActive }) => (isActive ? "nav-link active" : "nav-link")}
        >
          Pokédex
        </NavLink>
        <NavLink
          to="/favoritos"
          className={({ isActive }) => (isActive ? "nav-link active" : "nav-link")}
        >
          Favoritos
        </NavLink>
        <NavLink
          to="/adicionar"
          className={({ isActive }) => (isActive ? "nav-link active" : "nav-link")}
        >
          Adicionar
        </NavLink>
      </div>
    </nav>
  );
}

export default Navbar;