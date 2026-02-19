import Footer from "../user-geral/Footer";
import HeaderUser from "../user-geral/HeaderUser";
import Nav from "../user-geral/Nav";

const GroupHome = () => {
    return (
        <div>  
            <HeaderUser />
            <div className="div-nav-user">
                <Nav />
                <h1>main</h1>
            </div>
            <Footer />
        </div>
    )
}

export default GroupHome;