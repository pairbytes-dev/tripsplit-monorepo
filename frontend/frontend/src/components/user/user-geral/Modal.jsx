export default function Modal({ isOpen, setModalOpen}){

    if(isOpen){
    return (
        <div className="background-modal">
            <div className="main-div-modal">
                <section>
                    <div>
                        <h1>Criar Novo Grupo</h1>
                        <p>Informações do grupo</p>
                    </div>
                    <button onClick={setModalOpen}>Fechar</button>
                </section>
                <section>
                    <form action="get">
                        <label htmlFor="nome">Nome da despesa</label>
                        <input type="text" name="nome" id="nome"/>

                        <label htmlFor="categoria">Categoria</label>
                        <input type="range" name="categoria" id="categoria" />

                        <label htmlFor="descricao">Descrição</label>
                        <input type="text" name="descricao" id="descricao" />
                    </form>

                    <button onClick={setModalOpen}>Fechar</button>
                    <button onClick={setModalOpen}>Criar Grupo</button>
                </section>
            </div>
                
           
        </div>
    )
   }

   return null

}