import { useState } from 'react';

import './App.css'
import './style/Login.css';
import './style/Home.css';
import './style/MediaQuerie.css';

import Home from './components/Home';
import Login from './components/Login'
import SignIn from "./components/login/SignIn";
import SignUp from "./components/login/SignUp";

function App() {
  const [tela, setTela] = useState('home');

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
