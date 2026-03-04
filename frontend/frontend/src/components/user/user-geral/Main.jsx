import mais from '../../../assets/img/mais.png'

const Main = () => {
    return (
        <div className="user-main"> 
            <div className='user-main-div-title'>
                <div>
                    <h1 className="title-main-user">Meus Grupos</h1>
                    <p>Gerencie suas despesas compartilhadas</p>
                </div>
                <button className="btn-home btn-header" id='btn-green'>
                    <img src={mais} alt="icone de mais" className='icon-header'/>
                    Novo grupo
                </button>
            </div>
            <div>
                
            </div>
        </div>
    )
}

export default Main;