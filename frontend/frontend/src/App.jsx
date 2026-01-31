import { useState } from 'react';
import './App.css'
import './style/Login.css';
import './style/Home.css';

import Home from './components/Home';
import Login from './components/Login'
import SignIn from "./components/SignIn";
import SignUp from "./components/SignUp";

function App() {
  const [tela, setTela] = useState('home');
  console.log("Tela atual:", tela);

  return (
     <>
      {tela === 'home' && (
        <Home 
          irParaLogin={() => setTela("signin")}
          irParaCadastro={() => setTela("signup")}
        />
      )}

      {tela === 'signin' && (
        <Login irParaHome={() => setTela("home")}>
          <SignIn irParaCadastro={() => setTela("signup")} />
        </Login>
      )}

      {tela === "signup" && (
        <Login irParaHome={() => setTela("home")}>
          <SignUp irParaLogin={() => setTela("signin")} />
        </Login>
      )}
     </>
  )
}

export default App
