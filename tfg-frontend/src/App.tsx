import React, { useEffect, useState } from "react";
import { BrowserRouter, Route } from "react-router-dom";
import "./App.css";
import Nav from "./components/Nav";
import ShowForum from "./pages/ShowForum";
import Forums from "./pages/Forums";
import Home from "./pages/Home";
import Login from "./pages/Login";
import Register from "./pages/Register";

function App() {
    const [name, setName] = useState('');
    
    useEffect(() => {
        (
            async () => {
                const res = await fetch('http://127.0.0.1:8000/api/user',{
                    headers: {'Content-Type':'application/json'},
                    credentials: 'include',
                });

                const data = await res.json();

                setName(data.name)
            }
        )();
    },[setName]);
  return (
    <div className="App">
      <BrowserRouter>
        <Nav name={name} setName={setName} />
        <div className="container">
          <Route path="/" exact component={() => <Home name={name} setName={setName} />} />
          <Route path="/login" exact component={() => <Login setName={setName} />} />
          <Route path="/register" component={Register} />
          <Route path="/forums" component={Forums} />
          <Route path="/forum/:id" component={ShowForum} />
        </div>
      </BrowserRouter>
    </div>
  );
}

export default App;
