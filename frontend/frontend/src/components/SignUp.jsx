const SignUp = ({irParaLogin}) => {
    return(
        <div className="login-content">     
            <div className="login-title">
                <h1>Criar conta</h1>
                <p>Comece a dividir despesas agora</p>
            </div>
            <div className="login-form">
                <form className="form">
                    <label htmlFor="name">Nome completo</label>
                    <div className="input-wrapper">
                        <i class="fa-regular fa-user icon-input"></i>
                        <input type="text" name="name" id="name" placeholder="Digite seu nome completo"/>
                    </div>
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
                    <button type="submit" className="btn">Criar conta</button>
                </form>
            </div>
            <div className="login-text">
                <p>Já tem uma conta? <a href="#" onClick={irParaLogin}>Entrar</a></p>
            </div>
        </div>
    )
}

export default SignUp;