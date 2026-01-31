const SignIn = ({ irParaCadastro }) => {
    return(
        <div className="login-content">     
            <div className="login-title">
                <h1>Bem-vindo de volta!</h1>
                <p>Entre para acessar suas viagens</p>
            </div>
            <div className="login-form">
                <form className="form">
                    <label htmlFor="email">Email</label>
                    <div className='input-wrapper'>
                        <i class="fa-regular fa-envelope icon-input"></i>
                        <input type="email" name="email" id="email" placeholder="Digite seu email"/>
                    </div>
                    <label htmlFor="senha">Senha</label>
                    <div className="input-wrapper">
                        <i class="fa-solid fa-lock icon-input"></i>
                        <input type="password" name="senha" id="senha" placeholder="Digite sua senha"/>
                    </div>
                    <button type="submit" className="btn">Entrar</button>
                </form>
            </div>
            <div className="login-text">
                <p>Não tem uma conta?<a href="#" onClick={irParaCadastro}>Cadastre-se</a></p>
                <a href="#" className="password-text">Esqueceu sua senha?</a>
            </div>
        </div>
    )
}

export default SignIn;