import mais from '../../../assets/img/mais.png'
import grupo from '../../../assets/img/grupo.png'
import despesas from '../../../assets/img/despesas.png'
import planilha from '../../../assets/img/planilha.png'

import { useState } from 'react';
import Modal from './Modal.jsx';

const Nav = () => {

    const [openModal, setOpenModal] = useState(false);

    return (
        <div className="nav-div">
            <button className="btn-home btn-nav btn-header" onClick={() => setOpenModal(true)}>
                <img src={mais} alt="icone de mais" className='icon-header'/>
                Criar novo grupo
            </button>
            <ul className='ul-list'>
                <li className='li-list'>
                    <img src={grupo} alt="icone de grupo" className='icon-header'/>
                    Meus Grupos
                </li>
                <li className='li-list'>
                    <img src={despesas} alt="icone de gastos" className='icon-header'/>
                    Minhas despesas
                </li>
                <li className='li-list'>
                    <img src={planilha} alt="icone de planilha" className='icon-header'/>
                    Despesas do grupo
                </li>
            </ul>

            <Modal isOpen={openModal} setModalOpen={() => setOpenModal(!openModal)} />
        </div>
    )
}

export default Nav;