// src/services/jogosService.js
import api from './api';

// GET /jogos
export const listarJogos = async () => {
    const response = await api.get('/jogos');
    return response.data;
};

// GET /jogos/:id
export const buscarJogoPorID = async (id) => {
    const response = await api.get(`/jogos/${id}`);
    return response.data;
};

// POST /jogos
export const criarJogo = async (jogo) => {
    // Certifique-se que o objeto `jogo` tem os campos Nome, Descricao e Categoria
    const response = await api.post('/jogos', jogo);
    return response.data;
};

// PUT /jogos/:id
export const atualizarJogo = async (id, dados) => {
    // Certifique-se que o objeto `dados` tem os campos Nome, Descricao e Categoria
    const response = await api.put(`/jogos/${id}`, dados);
    return response.data;
};

// DELETE /jogos/:id
export const deletarJogo = async (id) => {
    const response = await api.delete(`/jogos/${id}`);
    return response.data; // Retorna a mensagem de sucesso
};