import Header from "./home/Header";
import HomeStart from "./home/HomeStart";

const Home = ({ irParaLogin, irParaCadastro }) => {
    return (
        <main className="home-main">
            <Header 
                irParaLogin={irParaLogin}
                irParaCadastro={irParaCadastro}
            />
            <HomeStart irParaCadastro={irParaCadastro}/>
        </main>
    )
}

export default Home;