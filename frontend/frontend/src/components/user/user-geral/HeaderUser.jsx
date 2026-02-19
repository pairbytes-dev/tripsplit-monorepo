import logoImg from '../../../assets/img/logo-icon.png';
import profile from '../../../assets/img/profile.png';
import exit from '../../../assets/img/exit.png';

const HeaderUser = () => {
    return (
        <header className="home-header">
            <div className='div-header'>
                <img src={logoImg} alt="Logo Racha.ai" className='logo-home'/> 
                <h1 className='title'>Racha<span>.ai</span></h1> 
            </div>
            <div className='div-header'>
                <button className='btn-home btn-header'>
                    <img src={profile} alt="icon de perfil" className='icon-header'/>
                    Nome do Usuário
                </button>
                <button className='btn-home btn-header'>
                    <img src={exit} alt="icon de saída" className='icon-header'/>
                    Sair
                </button>
            </div>
        </header>
    )
}

export default HeaderUser;