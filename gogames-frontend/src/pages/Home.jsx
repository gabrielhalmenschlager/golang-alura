// src/pages/Home.jsx
import React, { useEffect, useState } from 'react';
import { listarJogos } from '../services/jogosService'; // ajuste o caminho conforme teu projeto
import CardJogo from '../components/CardJogo';
import '../styles/Home.css';

const Home = () => {
    const [jogos, setJogos] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        let mounted = true;
        const fetchJogos = async () => {
            try {
                const data = await listarJogos();
                if (mounted) setJogos(data);
            } catch (err) {
                console.error('Erro ao carregar jogos:', err);
                if (mounted) setError('Não foi possível carregar os jogos.');
            } finally {
                if (mounted) setLoading(false);
            }
        };
        fetchJogos();
        return () => { mounted = false; };
    }, []);

    if (loading) {
        return (
            // Renderiza Loading no centro da viewport
            <div className="home-container loading-state">
                <h2 className="loading-text">CARREGANDO DADOS...</h2>
            </div>
        );
    }
    if (error) {
        return (
            <div className="home-container error-state">
                <h2 className="error-text">ERRO: {error}</h2>
            </div>
        );
    }

    return (
        <main className="home-container"> {/* Mudança de div para main (Semântica) */}
            <h1 className="main-title">
                {/* Remove o emoji 🕹️ do HTML para que o CSS de glitch funcione perfeitamente */}
                CATÁLOGO GOGAMES
            </h1>

            <section className="jogos-grid" aria-label="Lista de Jogos"> {/* section e aria-label para acessibilidade */}
                {jogos && jogos.length > 0 ? (
                    jogos.map(jogo => (
                        <CardJogo key={jogo.id} jogo={jogo} />
                    ))
                ) : (
                    // Adiciona um container para que a mensagem de "vazio" não se espalhe pelo grid
                    <div className="no-games-message-wrapper">
                        <p className="neon-message">
                            Nenhum jogo registrado. Adicione seu primeiro jogo!
                        </p>
                    </div>
                )}
            </section>
        </main>
    );
};

export default Home;