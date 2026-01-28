import imgLogo from '../assets/img/logo.png';
import imgArrow from '../assets/img/left-arrow.svg';
import background from '../assets/img/bg.jpeg';
import SignIn from './SignIn';
import SignUp from './SignUp';

const Login = () => {
    return(
        <main className="login-main">
            <section className="login-info"> 
                <div className="login-header">
                    <img src={imgLogo} alt="Logo" className='img-logo'/>
                    <button className='btn-link'>
                        <img src={imgArrow} alt="Voltar" />
                    </button>
                </div>     

                {/* <SignIn /> */}
                <SignUp />

            </section>
            <section className="login-img">
              <img src={background} alt="Imagem de fundo" className='background-img'/>
            </section>
        </main>
    )
};

export default Login;