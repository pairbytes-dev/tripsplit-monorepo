import { useState } from "react";

const SignIn = ({ irParaCadastro }) => {
    const [email, setEmail] = useState('');
    const [senha, setSenha] = useState('');

    const [mensagem, setMensagem] = useState("");
    const [tipo, setTipo] = useState("");

    const handleSubmit = async (e) => {
        e.preventDefault();

        try{
            const response = await fetch('http://localhost:8080/v1/auth/login', {
                method: 'POST',
                headers: {'Content-Type': 'application/json'},
                body: JSON.stringify({ email, password: senha }),
            });

            let data = null;

            try {
                data = await response.json();
            } catch {
                data = null;
            }

            if(response.status === 200){ //email e senha corretos

                if(data?.token){
                    localStorage.setItem("token", data.token);
                }
                
                setMensagem("Login realizado com sucesso!");
                setTipo("sucesso");

            } else if(response.status === 401){ //email ou senha incorretos
                
                setMensagem("Email ou senha incorretos");
                setTipo("erro");

            } else{
                setMensagem("Erro no login. Tente novamente.");
                setTipo("erro");
            }

        } catch(error){
            console.error("Erro no login:", error);
            setMensagem("Erro ao fazer login. Tente novamente.");
            setTipo("erro");
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

                {mensagem && (
                    <div className={`alert ${tipo}`}>
                    {mensagem}
                    <button onClick={() => {setMensagem(""); setTipo("");}} className="close-btn-alert">
                        X
                    </button>
                    </div>
                )}

                <p>Não tem uma conta?<span onClick={irParaCadastro}> Cadastre-se</span></p>
                <a href="#" className="password-text">Esqueceu sua senha?</a>
            </div>
        </div>
    )
}

export default SignIn;