import mais from '../../../assets/img/mais.png'
import { useState } from 'react';
import Modal from './Modal.jsx';

const Main = () => {

    const [openModal, setOpenModal] = useState(false);

    return (
        <div className="user-main"> 
            <div className='user-main-div-title'>
                <div>
                    <h1 className="title-main-user">Meus Grupos</h1>
                    <p>Gerencie suas despesas compartilhadas</p>
                </div>
                <button className="btn-home btn-header" id='btn-green' onClick={() => setOpenModal(true)}>
                    <img src={mais} alt="icone de mais" className='icon-header'/>
                    Novo grupo
                </button>
            </div>
            <div>
                
            </div>

            <Modal isOpen={openModal} setModalOpen={() => setOpenModal(!openModal)} />
        </div>
    )
}

export default Main;