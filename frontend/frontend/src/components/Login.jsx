const Login = () => {
    return(
        <div className="login-container">
            <div className="login-title">
                <h1>Bem-vindo de volta!</h1>
                <p>Entre para acessar suas viagens</p>
            </div>
            <div className="login-form">
                <form className="form">
                    <label htmlFor="email">Email</label>
                    <input type="email" name="email" id="email" placeholder="Digite seu email"/>
                    <label htmlFor="senha">Senha</label>
                    <input type="password" name="senha" id="senha" placeholder="Digite sua senha"/>
                    <button type="submit" className="btn">Entrar</button>
                </form>
            </div>
            <div className="login-text">
                <a href="#">Não tem uma conta? Cadastre-se</a>
                <a href="#">Esqueceu sua senha?</a>
            </div>
        </div>
    )
};

export default Login;