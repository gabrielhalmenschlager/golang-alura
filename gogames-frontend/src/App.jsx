// src/App.jsx
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Home from './pages/Home';
import CadastroJogo from './pages/CadastroJogo';
import DetalheJogo from './pages/DetalheJogo';
import Alugueis from './pages/Alugueis';
import Header from './components/Header';
import Footer from './components/Footer'; // Importe o Footer
import './styles/global.css';

function App() {
  return (
    <Router>
        {/* A div principal agora envelopa tudo para o layout flex completo */}
        <div className="app-container"> {/* Nova div para flex-grow */}
            <Header />
            <main className="container-neon"> {/* Ocupa o espaço restante */}
                <Routes>
                    <Route path="/" element={<Home />} /> 
                    <Route path="/jogos/:id" element={<DetalheJogo />} />
                    <Route path="/jogos/novo" element={<CadastroJogo />} />
                    <Route path="/alugueis" element={<Alugueis />} />
                </Routes>
            </main>
            <Footer /> {/* Adicione o Footer */}
        </div>
    </Router>
  );
}

export default App;