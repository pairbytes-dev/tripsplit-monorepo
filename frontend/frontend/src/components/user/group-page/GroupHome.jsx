import Footer from "../user-geral/Footer";
import HeaderUser from "../user-geral/HeaderUser";
import Nav from "../user-geral/Nav";
import Main from "../user-geral/Main";

const GroupHome = () => {
    return (
        <div className="div-group-user">   
            <HeaderUser />
            <div className="div-nav-user">  
                <Nav />
                <Main />
            </div>
            <Footer />
        </div>
    )
}

export default GroupHome;