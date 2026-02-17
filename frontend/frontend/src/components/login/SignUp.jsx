import React, { useState } from 'react';
import Modal from './Modal';

const SignUp = ({irParaLogin}) => {
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [senha, setSenha] = useState('');

    const [mensagem, setMensagem] = useState("");
    const [tipo, setTipo] = useState("");

    const tamSenha = senha.length >= 8;
    const hasNumber = /\d/.test(senha);
    const hasLetter = /[a-zA-Z]/.test(senha);

    const isSenhaValida = tamSenha && hasNumber && hasLetter;

    const handleSubmit = async (e) => {
        e.preventDefault();

        console.log('Enviando:', { name, email, password: senha });

        try{
            const response = await fetch('http://localhost:8080/v1/auth/register', {
                method: 'POST',
                headers: {'Content-Type': 'application/json'},
                body: JSON.stringify({ 
                    name,
                    email,
                    password: senha 
                })
            });

            console.log(response.status); //debugando

            if (!response.ok) { //verifica se deu erro
                const text = await response.text() || ""; //le o texto da resposta, que pode conter a mensagem de erro do servidor

                console.error('Erro server:', text || response.statusText);
                console.log(text || `Erro ${response.status}`);
                
                if(response.status === 409){ //email já existe
                    setMensagem("Este email já está cadastrado");
                    setTipo("erro");
                } else{
                    setMensagem("Erro ao cadastrar email.");
                    setTipo("erro");
                } 
                
                return;          
            }

            const data = await response.json();
            console.log(data);

            localStorage.setItem("token", data.token);
            setMensagem("Cadastro realizado com sucesso!");
            setTipo("sucesso");


        } catch(error){
            console.error("Erro no cadastro:", error);
            setMensagem("Erro ao fazer cadastro. Tente novamente.");
            setTipo("erro");
        }  
    }

    return(
        <div className="login-content"> 
            <div className="login-title">
                <h1>Criar conta</h1>
                <p>Comece a dividir despesas agora</p>
            </div>
            <div className="login-form">
                <form className="form" onSubmit={handleSubmit}>
                    <label htmlFor="name">Nome completo</label>
                    <div className="input-wrapper">
                        <i className="fa-regular fa-user icon-input"></i>
                        <input type="text" name="name" id="name" placeholder="Digite seu nome completo" 
                            value={name} onChange={(e) => setName(e.target.value)} required />
                    </div>
                    <label htmlFor="email">Email</label>
                    <div className='input-wrapper'>
                        <i className="fa-regular fa-envelope icon-input"></i>
                        <input type="email" name="email" id="email" placeholder="Digite seu email" 
                            value={email} onChange={(e) => setEmail(e.target.value)} required />
                    </div>
                    <label htmlFor="senha">Senha</label>
                    <div className="input-wrapper">
                        <i className="fa-solid fa-lock icon-input"></i>
                        <input type="password" name="senha" id="senha" placeholder="Digite sua senha" 
                            value={senha} onChange={(e) => setSenha(e.target.value)} required />
                    </div>  
                        {senha.length > 0 && !isSenhaValida && (
                            <p className="error-message">
                            A senha precisa ter pelo menos 8 caracteres incluindo letras e números.
                            </p>
                        )}
                    <button type="submit" className="btn">Criar conta</button>
                </form>
            </div>
            <div className="login-text">

                {mensagem && (
                    <div className={`alert ${tipo}`}>
                    {mensagem}
                    <button onClick={() => {setMensagem(""); setTipo("");}} className="close-btn-alert">
                        X
                    </button>
                    </div>
                )}

                <p>Já tem uma conta?<span onClick={irParaLogin}> Entrar</span></p>
            </div>

        </div>
    )
}

export default SignUp;