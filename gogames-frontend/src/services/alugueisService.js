// src/services/alugueisService.js
import api from './api';

// GET /alugueis
export const listarAlugueis = async () => {
    const response = await api.get('/alugueis');
    return response.data;
};

// POST /alugueis/alugar/:id (O ID aqui é o JogoID)
export const alugarJogo = async (jogoId) => {
    // Sua API espera o ID na URL e não um corpo JSON, conforme o controlador:
    // r.POST("/alugueis/alugar/:id", controllers.AlugarJogo)
    const response = await api.post(`/alugueis/alugar/${jogoId}`);
    return response.data;
};

// PUT /alugueis/devolver/:id (O ID aqui é o AluguelID, que será buscado na tela de Aluguéis)
export const devolverJogo = async (aluguelId) => {
    const response = await api.put(`/alugueis/devolver/${aluguelId}`);
    return response.data;
};

// GET /alugueis/ativos
export const listarAlugueisAtivos = async () => {
    const response = await api.get('/alugueis/ativos');
    return response.data;
};

// GET /alugueis/inativos
export const listarAlugueisInativos = async () => {
    const response = await api.get('/alugueis/inativos');
    return response.data;
};