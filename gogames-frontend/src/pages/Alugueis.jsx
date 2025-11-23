// src/pages/Alugueis.jsx (Versão COMPLETA e CORRIGIDA)

import React, { useState, useEffect } from 'react';
import { listarAlugueis, listarAlugueisAtivos, listarAlugueisInativos, devolverJogo } from '../services/alugueisService';
import { RefreshCw } from 'lucide-react'; // Ícone de recarregar (Verifique se 'lucide-react' está instalado!)
import '../styles/Alugueis.css';

const Alugueis = () => {
    const [alugueis, setAlugueis] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const [filter, setFilter] = useState('TODOS'); // TODOS, ALUGADO, DEVOLVIDO
    const [message, setMessage] = useState('');

    const fetchAlugueis = async (currentFilter) => {
        setLoading(true);
        setError(null);
        setMessage('');

        let serviceCall;
        if (currentFilter === 'ALUGADO') {
            serviceCall = listarAlugueisAtivos;
        } else if (currentFilter === 'DEVOLVIDO') {
            serviceCall = listarAlugueisInativos;
        } else {
            serviceCall = listarAlugueis;
        }

        try {
            const data = await serviceCall();
            setAlugueis(data);
        } catch (err) {
            console.error('Erro ao carregar aluguéis:', err);
            // Verifica se a API está inacessível
            const errorMessage = err.message.includes('Network Error') 
                ? '⚠️ API Indisponível. Não foi possível carregar os dados.'
                : '❌ Não foi possível carregar os dados de aluguéis.';
            setError(errorMessage);
            setAlugueis([]); // Limpa a lista em caso de erro
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        fetchAlugueis(filter);
    }, [filter]);

    const handleDevolucao = async (aluguelID) => {
        if (!window.confirm("Confirma a devolução deste aluguel?")) {
            return;
        }

        try {
            await devolverJogo(aluguelID);
            setMessage(`✅ Aluguel #${aluguelID} devolvido com sucesso!`);
            // Recarrega a lista para atualizar o status e remover do filtro 'ALUGADO'
            fetchAlugueis(filter); 
        } catch (error) {
            console.error('Erro ao devolver:', error.response ? error.response.data.erro : error.message);
            setMessage(`❌ Erro na devolução: ${error.response ? error.response.data.erro : 'Falha na conexão'}`);
        }
    };

    const formatDate = (dateStr) => {
        if (!dateStr) return 'N/A';
        return new Date(dateStr).toLocaleDateString('pt-BR');
    }

    return (
        <div className="alugueis-container">
            {/* Título com estilo Press Start 2P e Glitch/Neon Pink */}
            <h1 className="main-title alugueis-arcade-title">ALUGUÉIS DA GO GAMES</h1>
            
            {message && (
                <p className={`feedback-message ${message.includes('Erro') || message.includes('❌') ? 'error' : 'success'}`}>
                    {message}
                </p>
            )}

            {/* Componente de Filtro (Abas) */}
            <div className="filter-tabs">
                <button 
                    onClick={() => setFilter('TODOS')}
                    className={`neon-tab ${filter === 'TODOS' ? 'active-cyan' : ''}`}
                >
                    TODOS ({alugueis.length})
                </button>
                <button 
                    onClick={() => setFilter('ALUGADO')}
                    className={`neon-tab ${filter === 'ALUGADO' ? 'active-red' : ''}`}
                >
                    ATIVOS
                </button>
                <button 
                    onClick={() => setFilter('DEVOLVIDO')}
                    className={`neon-tab ${filter === 'DEVOLVIDO' ? 'active-green' : ''}`}
                >
                    INATIVOS
                </button>
                {/* Botão de recarregar para UX */}
                <button 
                    onClick={() => fetchAlugueis(filter)}
                    className="neon-tab refresh-button"
                    disabled={loading}
                >
                    <RefreshCw size={16} className={loading ? 'spin' : ''} />
                </button>
            </div>

            {loading ? (
                <h2 className="loading-text">CARREGANDO REGISTROS...</h2>
            ) : error ? (
                <h2 className="error-text">{error}</h2>
            ) : alugueis.length === 0 ? (
                <p className="neon-message">Nenhum aluguel encontrado no filtro **{filter}**.</p>
            ) : (
                /* Tabela de Aluguéis */
                <div className="alugueis-table-wrapper">
                    <table className="neon-table">
                        <thead>
                            <tr>
                                <th>ID</th>
                                <th>JOGO</th>
                                <th>DATA ALUGUEL</th>
                                <th>DEVOL. PREVISTA</th>
                                <th>STATUS</th>
                                <th>AÇÕES / DATA REAL</th>
                            </tr>
                        </thead>
                        <tbody>
                            {alugueis.map(item => (
                                <tr key={item.id} className={item.status === 'ALUGADO' ? 'active-row' : 'inactive-row'}>
                                    <td>{item.id}</td>
                                    <td>{item.jogo.nome}</td>
                                    <td>{formatDate(item.data_aluguel)}</td>
                                    <td>{formatDate(item.data_devolucao_prevista)}</td>
                                    <td>
                                        <span className={`status-pill status-${item.status.toLowerCase()}`}>
                                            {item.status}
                                        </span>
                                    </td>
                                    <td>
                                        {item.status === 'ALUGADO' ? (
                                            <button 
                                                className="neon-button-small btn-devolver"
                                                onClick={() => handleDevolucao(item.id)}
                                            >
                                                DEVOLVER
                                            </button>
                                        ) : (
                                            <span className="devolvido-date">{formatDate(item.data_devolucao_real)}</span>
                                        )}
                                    </td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                </div>
            )}
        </div>
    );
};

export default Alugueis;