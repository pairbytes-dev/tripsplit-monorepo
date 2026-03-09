export default function Modal({ isOpen, setModalOpen}){

    if(isOpen){
    return (
        <div className="background-modal">
            <div className="main-div-modal">
                <section className="section-modal">
                    <div>
                        <h1 className="title-modal">Criar Novo Grupo</h1>
                        <p>Informações do grupo</p>
                    </div>
                    <button onClick={setModalOpen} className="btn-modal-close"><i class="fa-solid fa-x"></i></button>
                </section>
                <section className="section-modal-2">
                    <form action="get" className="form-modal">
                        <label htmlFor="nome">Nome da despesa</label>
                        <div className="div-modal-form"><input type="text" name="nome" id="nome" placeholder="Ex: Viagem Ubatuba 2026"/></div>

                        <label htmlFor="categoria">Categoria</label>
                        <div className="div-modal-form">
                            <select name="categoria" id="categoria">
                                <option value="">Selecione uma categoria</option>
                                <option value="alimentacao">Alimentação</option>
                                <option value="transporte">Transporte</option>
                                <option value="viagem">Viagem</option>
                                <option value="lazer">Lazer</option>
                                <option value="saude">Saúde</option>
                                <option value="outros">Outros</option>
                            </select>
                        </div>

                        <label htmlFor="descricao">Descrição</label>
                        <div className="div-modal-form"><textarea name="descricao" id="descricao"></textarea></div>
                    </form>

                    <div className="div-btn-modal">
                        <button onClick={setModalOpen} className="btn-home btn-close btn-size">Fechar</button>
                        <button onClick={setModalOpen} className="btn-home btn-size">Criar Grupo</button>
                    </div>
                </section>
            </div>
                
           
        </div>
    )
   }

   return null

}