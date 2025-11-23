// src/pages/DetalheJogo.jsx
import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { buscarJogoPorID, atualizarJogo, deletarJogo } from '../services/jogosService';
import { alugarJogo } from '../services/alugueisService';
import '../styles/DetalheJogo.css';
import { ArrowLeft } from 'lucide-react';

// Lista de categorias para o formulário de edição
const categorias = ["BRONZE", "PRATA", "OURO"];

const DetalheJogo = () => {
    const { id } = useParams(); // Pega o ID da URL
    const navigate = useNavigate();

    const [jogo, setJogo] = useState(null);
    const [isEditing, setIsEditing] = useState(false);
    const [formData, setFormData] = useState({});
    const [loading, setLoading] = useState(true);
    const [message, setMessage] = useState('');

    // --- 1. Fetch de Dados ---
    const fetchJogo = async () => {
        try {
            setLoading(true);
            const data = await buscarJogoPorID(id);
            setJogo(data);
            setFormData({
                nome: data.nome,
                descricao: data.descricao,
                categoria: data.categoria,
            });
            setLoading(false);
        } catch (error) {
            console.error('Erro ao buscar jogo:', error);
            setMessage('Jogo não encontrado ou erro de conexão.');
            setLoading(false);
        }
    };

    useEffect(() => {
        fetchJogo();
    }, [id]);

    // --- 2. Handlers de Edição ---
    const handleChange = (e) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
    };

    const handleUpdate = async (e) => {
        e.preventDefault();
        try {
            const updatedJogo = await atualizarJogo(id, formData);
            setJogo(updatedJogo);
            setIsEditing(false);
            setMessage('✨ Jogo atualizado com sucesso!');
        } catch (error) {
            console.error('Erro ao atualizar:', error);
            setMessage('🚨 Erro ao atualizar o jogo.');
        }
    };

    // --- 3. Handler de Aluguel ---
    const handleAlugar = async () => {
        if (!window.confirm(`Tem certeza que deseja alugar o jogo "${jogo.nome}"?`)) {
            return;
        }
        try {
            await alugarJogo(jogo.id);
            setMessage('🎉 Jogo ALUGADO com sucesso! Status atualizado.');
            // Atualiza o estado local para refletir a mudança
            setJogo(prev => ({ ...prev, disponivel: false }));
        } catch (error) {
            console.error('Erro ao alugar:', error.response ? error.response.data.erro : error.message);
            setMessage(`❌ Erro ao alugar: ${error.response ? error.response.data.erro : 'Serviço indisponível'}`);
        }
    };

    // --- 4. Handler de Exclusão ---
    const handleDeletar = async () => {
        if (!window.confirm(`AVISO! Confirma a exclusão de "${jogo.nome}"?`)) {
            return;
        }
        try {
            await deletarJogo(id);
            setMessage('🗑️ Jogo excluído com sucesso! Redirecionando...');
            setTimeout(() => navigate('/'), 2000); // Volta para a lista principal
        } catch (error) {
            console.error('Erro ao deletar:', error.response ? error.response.data.erro : error.message);
            setMessage(`❌ Não foi possível excluir: ${error.response ? error.response.data.erro : 'Verifique se o jogo está alugado.'}`);
        }
    };

    if (loading) return <h2 className="loading-text">BUSCANDO DADOS ARCADE...</h2>;
    if (!jogo) return <h2 className="error-text">{message}</h2>;

    return (
        <main className="detail-page-container container-neon"> {/* container-neon para padronização */}
            <button className="neon-button back-button" onClick={() => navigate('/')}>
                <ArrowLeft size={16} style={{ marginRight: '8px' }} />
                VOLTAR AO CATÁLOGO
            </button>

            <h1 className="detail-title">{jogo.nome}</h1>

            {message && <p className={`feedback-message ${message.includes('Erro') || message.includes('Não foi possível excluir') ? 'error' : 'success'}`}>{message}</p>}
            
            {!isEditing ? (
                // --- MODO VISUALIZAÇÃO: COM IMAGEM E INFORMAÇÕES LADO A LADO ---
                <div className="game-detail-content">
                    {/* COLUNA 1: IMAGEM/ARTE */}
                    <div className="game-image-wrapper">
                        {jogo.imagem_url ? (
                            <img src={jogo.imagem_url} alt={`Arte de ${jogo.nome}`} className="game-main-image" />
                        ) : (
                            <div className="game-placeholder-image">
                                <span>[ARTE NÃO ENCONTRADA]</span>
                            </div>
                        )}
                        <div className="image-scan-line" aria-hidden="true"></div> {/* Novo efeito de scanline */}
                    </div>

                    {/* COLUNA 2: DADOS E AÇÕES */}
                    <div className="game-info-actions neon-data-panel"> {/* Novo nome de classe */}
                        <div className="data-row status-info">
                            <h2 className="status-label">STATUS ATUAL:</h2>
                            <span className={`status-tag status-${jogo.disponivel ? 'available' : 'rented'}`}>
                                {jogo.disponivel ? 'DISPONÍVEL NO ARCADE' : 'ALUGADO (OFFLINE)'}
                            </span>
                        </div>
                        
                        <div className="data-row info-grid">
                            <p className="label-data">DESCRIÇÃO:</p>
                            <p className="value-data description-text">{jogo.descricao || 'Nenhuma descrição fornecida.'}</p>
                            
                            <p className="label-data">CATEGORIA:</p>
                            <span className={`category-tag category-${jogo.categoria}`}>{jogo.categoria}</span>
                            
                            <p className="label-data">DATA DE CADASTRO:</p>
                            <span className="value-data">{new Date(jogo.data_cadastro).toLocaleDateString()}</span>
                        </div>

                        <div className="action-buttons">
                            {/* Botão de Alugar */}
                            <button
                                className={`neon-button ${jogo.disponivel ? 'btn-green' : 'btn-disabled'}`}
                                onClick={handleAlugar}
                                disabled={!jogo.disponivel}
                            >
                                {jogo.disponivel ? 'ALUGAR' : 'INDISPONÍVEL'}
                            </button>

                            {/* Botões de Edição e Exclusão */}
                            <button className="neon-button btn-cyan" onClick={() => setIsEditing(true)}>
                                EDITAR REGISTRO
                            </button>
                            <button
                                className="neon-button btn-red"
                                onClick={handleDeletar}
                                disabled={!jogo.disponivel}
                            >
                                EXCLUIR REGISTRO
                            </button>
                        </div>
                    </div>
                </div>
            ) : (
                // --- MODO EDIÇÃO: FORMULÁRIO DE TERMINAL ---
                <form onSubmit={handleUpdate} className="edit-form neon-data-panel">
                    <h2>INTERFACE DE EDIÇÃO DE DADOS</h2>
                    <label className="neon-label">
                        NOME DO JOGO:
                        <input type="text" name="nome" value={formData.nome || ''} onChange={handleChange} required className="neon-input" />
                    </label>
                    <label className="neon-label">
                        DESCRIÇÃO:
                        <textarea name="descricao" value={formData.descricao || ''} onChange={handleChange} required className="neon-input"></textarea>
                    </label>
                    <label className="neon-label">
                        CATEGORIA:
                        <select name="categoria" value={formData.categoria || ''} onChange={handleChange} required className="neon-input">
                            {categorias.map(cat => <option key={cat} value={cat}>{cat}</option>)}
                        </select>
                    </label>
                    <div className="action-buttons edit-mode-buttons">
                        <button type="submit" className="neon-button btn-green">SALVAR ALTERAÇÕES</button>
                        <button type="button" className="neon-button btn-red" onClick={() => setIsEditing(false)}>CANCELAR EDIÇÃO</button>
                    </div>
                </form>
            )}
        </main>
    );
};

export default DetalheJogo;