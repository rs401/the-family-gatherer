import React from 'react'
import { Link } from 'react-router-dom'

const Nav = (props: {name: string, setName: (name: string) => void}) => {
    const logout = async () => {
        await fetch('http://127.0.0.1:8000/api/logout',{
            method: 'POST',
            headers: {'Content-Type':'application/json'},
            credentials: 'include',
        });
        props.setName('');
    }
    
    let menu;
    console.log('====== name in nav: ' + props.name)
    if (props.name === '' || props.name === undefined){
        menu = (
            <ul className="navbar-nav me-auto mb-2 mb-md-0">
            <li className="nav-item">
                <Link className="nav-link" to="/login">
                Login
                </Link>
            </li>
            <li className="nav-item">
                <Link className="nav-link" to="/register">
                Register
                </Link>
            </li>
            </ul>
        )
    } else {
        menu = (
            <ul className="navbar-nav me-auto mb-2 mb-md-0">
            <li className="nav-item">
              <Link className="nav-link" to="/login">
                Hello {props.name}
              </Link>
            </li>
            <li className="nav-item">
              <Link className="nav-link" to="/login" onClick={logout}>
                Logout
              </Link>
            </li>
            </ul>
        )
    }
    return (
        <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
        <div className="container-fluid">
          <Link to="/" className="navbar-brand">
            The Family Gatherer
          </Link>
          <button
            className="navbar-toggler"
            type="button"
            data-bs-toggle="collapse"
            data-bs-target="#navbarCollapse"
            aria-controls="navbarCollapse"
            aria-expanded="false"
            aria-label="Toggle navigation"
          >
            <span className="navbar-toggler-icon"></span>
          </button>
          <div className="collapse navbar-collapse" id="navbarCollapse">
            <ul className="navbar-nav me-auto mb-2 mb-md-0">
              <li className="nav-item">
                <Link className="nav-link active" aria-current="page" to="/">
                  Home
                </Link>
              </li>
              <li className="nav-item">
                <Link className="nav-link" to="#">
                  About
                </Link>
              </li>
              <li className="nav-item">
                <Link className="nav-link disabled" to="#">
                  Disabled
                </Link>
              </li>
            </ul>
            <div>
                { menu }
            </div>

          </div>
        </div>
      </nav>
    )
}

export default Nav
