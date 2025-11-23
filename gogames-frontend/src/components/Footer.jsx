// src/components/Footer.jsx (Versão com Ícones e Minimalista)
import React from 'react';
import { Github, Linkedin } from 'lucide-react'; // Importando ícones
import '../styles/Footer.css'; 

const Footer = () => {
  return (
    <footer className="neon-footer">
      <div className="footer-content">
        {/* Informação principal em uma linha */}
        <p>
          &copy; 2024 **GoGames API**. Todos os direitos reservados.
        </p>
        
        <div className="social-links">
          {/* Link desenvolvido com */}
          <span className="developed-by">
            Desenvolvido com <span className="neon-heart">❤️</span> e Go & React
          </span>

          {/* Links Sociais com Ícones */}
          <a 
            href="https://linkedin.com/in/seu-perfil" 
            target="_blank" 
            rel="noopener noreferrer" 
            className="neon-link-footer"
            title="LinkedIn">
            {/* O ícone será estilizado com a cor neon */}
            <Linkedin size={20} /> 
          </a>
          <a 
            href="https://github.com/seu-usuario" 
            target="_blank" 
            rel="noopener noreferrer" 
            className="neon-link-footer"
            title="GitHub">
            <Github size={20} />
          </a>
        </div>
      </div>
    </footer>
  );
};

export default Footer;