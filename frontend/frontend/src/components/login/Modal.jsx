const Modal = ({ onClose }) => {
    return(
        <div className="modal-container">
            <div className="modal">
                <h2>Cadastro realizado com sucesso! 🎉</h2>
                <button onClick={onClose}>X</button>
            </div>
        </div>
    )
}

export default Modal;