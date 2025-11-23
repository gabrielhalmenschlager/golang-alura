// src/pages/CadastroJogo.jsx (Revisado)

import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { criarJogo } from '../services/jogosService';
import { ArrowLeft } from 'lucide-react'; // Ícone para o botão de voltar
import '../styles/Formulario.css'; 

const categorias = ["BRONZE", "PRATA", "OURO"];

const CadastroJogo = () => {
    const navigate = useNavigate();
    const [formData, setFormData] = useState({
        nome: '',
        descricao: '',
        categoria: 'BRONZE',
        // Adicionando o campo de URL da Imagem
        imagem_url: '', 
    });
    const [loading, setLoading] = useState(false);
    const [message, setMessage] = useState('');

    const handleChange = (e) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        setLoading(true);
        setMessage('');

        try {
            await criarJogo(formData); 
            setMessage('🎉 Jogo cadastrado com sucesso! Redirecionando...');
            
            setTimeout(() => navigate('/'), 2000); 

        } catch (error) {
            console.error('Erro ao cadastrar:', error);
            setMessage('❌ Erro ao cadastrar o jogo. Verifique os dados e a API.');
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="form-container">
            {/* Título com Press Start 2P e Glitch, alinhado com Home.css */}
            <h1 className="main-title form-arcade-title">NOVO JOGO</h1>
            
            {message && (
                <p className={`feedback-message ${message.includes('Erro') ? 'error' : 'success'}`}>
                    {message}
                </p>
            )}

            <form onSubmit={handleSubmit} className="neon-form">
                
                <label className="neon-label">
                    Nome do Jogo:
                    <input 
                        type="text" 
                        name="nome" 
                        value={formData.nome} 
                        onChange={handleChange} 
                        required 
                        className="neon-input" 
                        disabled={loading}
                    />
                </label>
                
                <label className="neon-label">
                    Descrição:
                    <textarea 
                        name="descricao" 
                        value={formData.descricao} 
                        onChange={handleChange} 
                        required 
                        className="neon-input textarea-neon" 
                        disabled={loading}
                    ></textarea>
                </label>

                <label className="neon-label">
                    URL da Imagem (Opcional):
                    <input 
                        type="url" 
                        name="imagem_url" 
                        value={formData.imagem_url} 
                        onChange={handleChange} 
                        className="neon-input" 
                        disabled={loading}
                        placeholder="https://exemplo.com/capa.jpg"
                    />
                </label>
                
                <label className="neon-label">
                    Categoria:
                    <select 
                        name="categoria" 
                        value={formData.categoria} 
                        onChange={handleChange} 
                        required 
                        className="neon-input select-neon"
                        disabled={loading}
                    >
                        {categorias.map(cat => (
                            <option key={cat} value={cat}>{cat}</option>
                        ))}
                    </select>
                </label>
                
                <button 
                    type="submit" 
                    // Reutiliza o estilo global 'neon-button'
                    className="neon-button btn-submit-green"
                    disabled={loading}
                >
                    {loading ? 'SALVANDO...' : 'CADASTRAR JOGO'}
                </button>
            </form>

            <button className="neon-button back-button back-to-home" onClick={() => navigate('/')} disabled={loading}>
                <ArrowLeft size={16} /> CANCELAR / VOLTAR
            </button>
        </div>
    );
};

export default CadastroJogo;