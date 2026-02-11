import items from "../../assets/javascript/script";

const HomeStart = ({ irParaCadastro }) => {
    return (
      <div className="div-home-start">
        <section className="section-start">
          <p className="first-txt-home"> {/* <-- ALTERAR --> */}
            Comece a dividir suas despesas. Organize sem dor de cabeça na hora de
            dividir.
          </p>
          <p className="second-txt-home"> {/* <-- ALTERAR --> */}
            Divida gastos de forma justa, automática e
            transparente. Registre despesas, veja relatórios em tempo real e
            acerte as contas em segundos.
          </p>
          <button className="btn-home" id="start-now" onClick={irParaCadastro}>Começar a dividir</button> 
        </section>
        <section className="section-start">
          <div className="expenses-start">
            <div className="expenses-div-header"> {/* <-- ALTERAR --> */}
                <div className="expenses-div-txt">
                  <p>Resumo de despesas</p>
                  <p className="gray-text">3 despesas</p>
                </div>
                <div className="expenses-div-txt">
                  <p className="gray-text">Total das despesas </p> 
                  <p id="valor">R$ 3.420,00</p>
                </div>
            </div>
            <div className="expenses-div-items-container">
                {items.map(item => (
                    <div key={item.id} className="expenses-div-items">
                        <div className="items-icon">{item.icon}</div>
                        <p>{item.name}</p>
                        <p className="items-price">{item.price}</p>
                    </div>
                ))}
            </div>
            <hr />
            <div className="expenses-div-header" id="final-value"> 
              <p>Você deve</p>
              <p>R$855,00</p> {/* <-- ALTERAR --> */}
            </div>
          </div>
        </section>
      </div>
    );
}

export default HomeStart;