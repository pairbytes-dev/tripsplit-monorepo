import { useState } from "react";

const SignIn = ({ irParaCadastro }) => {
    const [email, setEmail] = useState('');
    const [senha, setSenha] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();

        try{
            const response = await fetch('http://localhost:8080/v1/auth/login', {
                method: 'POST',
                headers: {'Content-Type': 'application/json'},
                body: JSON.stringify({ email, password: senha }),
            });

            const data = await response.json();

            if(!response.ok){
                alert(data.message || "Erro ao fazer login");
                return;
            }

            localStorage.setItem("token", data.token);

            alert("Login bem-sucedido!");
        } catch(error){
            console.error("Erro no login:", error);
            alert("Não foi possível conectar ao servidor. Tente novamente.");
        }  
    }

    return(
        <div className="login-content">     
            <div className="login-title">
                <h1>Bem-vindo de volta!</h1>
                <p>Entre para começar a dividir</p>
            </div>
            <div className="login-form">
                <form className="form" onSubmit={handleSubmit}>
                    <label htmlFor="email">Email</label>
                    <div className='input-wrapper'>
                        <i className="fa-regular fa-envelope icon-input"></i>
                        <input 
                            type="email" 
                            name="email" 
                            id="email" 
                            placeholder="Digite seu email" 
                            value={email}
                            onChange={(e) => setEmail(e.target.value)}
                            required 
                        />
                    </div>
                    <label htmlFor="senha">Senha</label>
                    <div className="input-wrapper">
                        <i className="fa-solid fa-lock icon-input"></i>
                        <input 
                            type="password" 
                            name="senha" 
                            id="senha" 
                            placeholder="Digite sua senha" 
                            value={senha}
                            onChange={(e) => setSenha(e.target.value)}
                            required/>
                    </div>
                    <button type="submit" className="btn">Entrar</button>
                </form>
            </div>
            <div className="login-text">
                <p>Não tem uma conta?<span onClick={irParaCadastro}> Cadastre-se</span></p>
                <a href="#" className="password-text">Esqueceu sua senha?</a>
            </div>
        </div>
    )
}

export default SignIn;