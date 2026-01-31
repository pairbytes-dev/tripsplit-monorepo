import items from "../../assets/javascript/script";

const HomeStart = ({ irParaCadastro }) => {
    return (
      <div className="div-home-start">
        <section className="section-start">
          <p className="first-txt-home">
            Comece a dividir suas despesas. Viaje sem dor de cabeça na hora de
            dividir.
          </p>
          <p className="second-txt-home">
            Divida despesas de viagens de forma justa, automática e
            transparente. Registre gastos, veja relatórios em tempo real e
            acerte as contas em segundos.
          </p>
          <button className="btn-home" id="start-now" onClick={irParaCadastro}>Começar agora</button>
        </section>
        <section className="section-start">
          <div className="expenses-start">
            <div className="expenses-div-header">
                <div className="expenses-div-txt">
                  <p>Viagem - Rio de Janeiro</p>
                  <p className="gray-text">4 pessoas</p>
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
              <p>R$855,00</p>
            </div>
          </div>
        </section>
      </div>
    );
}

export default HomeStart;