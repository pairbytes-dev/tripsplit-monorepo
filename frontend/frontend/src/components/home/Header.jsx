import imgLogo from '../../assets/logo-icon.png';

const Header = () => {
    return (
        <header className='home-header'>
            <div className='div-header'>
                <img src={imgLogo} alt="Logo TripSplit" className='logo-home'/>
                <h1 className='title'>Trip<span>Split</span></h1>
            </div>
            <div className='div-header'>
                <button className='btn-home' id='btn-login'>Login</button>
                <button className='btn-home'>Cadastre-se</button>
            </div>
        </header>
    )
}   

export default Header;