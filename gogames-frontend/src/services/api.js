// src/services/api.js
import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost:8000', // Certifique-se que é a porta correta da sua API Go
  headers: {
    'Content-Type': 'application/json',
  },
});

export default api;