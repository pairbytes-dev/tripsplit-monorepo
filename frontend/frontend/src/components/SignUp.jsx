import React, { useState } from 'react';

const SignUp = ({irParaLogin}) => {
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [senha, setSenha] = useState('');
    
    const handleSubmit = async (e) => {
        e.preventDefault();

        console.log('Enviando:', { name, email, password: senha }); //debugando

        try{
            const response = await fetch('http://localhost:8080/v1/auth/register', {
                method: 'POST',
                headers: {'Content-Type': 'application/json'},
                body: JSON.stringify({ 
                    name,
                    email,
                    password: senha 
                }),
            });

            console.log(response.status); //debugando

            if (!response.ok) {
            const text = await response.text();
            console.error('Erro server:', text || response.statusText);
            alert(text || `Erro ${response.status}`);
            return;
        }

            const data = await response.json();
            console.log(data);

            localStorage.setItem("token", data.token);
            alert("Cadastro bem-sucedido!");

        } catch(error){
            console.error("Erro no cadastro:", error);
            alert("Não foi possível conectar ao servidor. Tente novamente.");
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
                        <input type="text" name="name" id="name" placeholder="Digite seu nome completo" value={name}
                            onChange={(e) => setName(e.target.value)} required />
                    </div>
                    <label htmlFor="email">Email</label>
                    <div className='input-wrapper'>
                        <i className="fa-regular fa-envelope icon-input"></i>
                        <input type="email" name="email" id="email" placeholder="Digite seu email" value={email}
                            onChange={(e) => setEmail(e.target.value)} required />
                    </div>
                    <label htmlFor="senha">Senha</label>
                    <div className="input-wrapper">
                        <i className="fa-solid fa-lock icon-input"></i>
                        <input type="password" name="senha" id="senha" placeholder="Digite sua senha" value={senha}
                            onChange={(e) => setSenha(e.target.value)} required/>
                    </div>  
                    <button type="submit" className="btn">Criar conta</button>
                </form>
            </div>
            <div className="login-text">
                <p>Já tem uma conta?<span onClick={irParaLogin}> Entrar</span></p>
            </div>
        </div>
    )
}

export default SignUp;