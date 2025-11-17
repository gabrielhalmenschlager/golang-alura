import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';

function EditPokemon() {
    const { id } = useParams();
    const navigate = useNavigate();
    const [pokemon, setPokemon] = useState({ name: "", type: "", image_url: "" });
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        fetch(`http://localhost:8000/pokemon/${id}`)
            .then(response => {
                if (!response.ok) throw new Error('Pokémon não encontrado');
                return response.json();
            })
            .then(data => {
                setPokemon(data);
                setLoading(false);
            })
            .catch(error => {
                console.error("Erro ao carregar Pokémon:", error);
                setLoading(false);
            });
    }, [id]);

    const handleChange = (e) => {
        setPokemon({ ...pokemon, [e.target.name]: e.target.value });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        
        await fetch(`http://localhost:8000/pokemon/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(pokemon),
        });

        alert(`Pokémon #${id} atualizado com sucesso!`);
        navigate('/');
    };

    if (loading) {
        return <div className="loading-text">Carregando dados do Pokémon...</div>;
    }

    return (
        <div className="form-container">
            <form onSubmit={handleSubmit} className="pokedex-form">
                <h2 className="page-title">Editar Pokémon #{id}</h2>
                <div className="form-group">
                    <label htmlFor="name">Nome:</label>
                    <input id="name" name="name" placeholder="Nome" value={pokemon.name} onChange={handleChange} required />
                </div>
                <div className="form-group">
                    <label htmlFor="type">Tipo:</label>
                    <input id="type" name="type" placeholder="Tipo" value={pokemon.type} onChange={handleChange} required />
                </div>
                <div className="form-group">
                    <label htmlFor="image_url">URL da Imagem:</label>
                    <input id="image_url" name="image_url" placeholder="URL da imagem" value={pokemon.image_url} onChange={handleChange} required />
                </div>
                <button type="submit" className="btn submit-btn" style={{backgroundColor: '#007bff'}}>
                    Salvar Alterações
                </button>
            </form>
        </div>
    );
}

export default EditPokemon;