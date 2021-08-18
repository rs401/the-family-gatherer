import React, { SyntheticEvent, useState } from 'react'
import { Redirect } from 'react-router-dom';

const Login = (props: {setName: (name: string) => void}) => {
    const [email, setEmail] = useState('');
    const [password, setPass] = useState('');
    const [redirect, setRedirect] = useState(false);

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();
        const response = await fetch('http://127.0.0.1:8000/api/login',{
            method: 'POST',
            headers: {'Content-Type':'application/json'},
            credentials: 'include',
            body: JSON.stringify({
                email,
                password
            })
        });

        const data = await response.json();
        setRedirect(true);
        props.setName(data.name)
    }

    if(redirect){
        return <Redirect to="/" />
    }

    return (
        <div className="form-signin">
            <form onSubmit={submit}>
            <h1 className="h3 mb-3 fw-normal">Please sign in</h1>

            <div className="form-floating">
                <input
                type="email"
                className="form-control"
                id="floatingInput"
                name="email"
                placeholder="name@example.com"
                onChange={e => setEmail(e.target.value)}
                />
                <label htmlFor="floatingInput">Email</label>
            </div>
            <div className="form-floating">
                <input
                type="password"
                className="form-control"
                id="floatingPassword"
                name="password"
                placeholder="Password"
                onChange={e => setPass(e.target.value)}
                />
                <label htmlFor="floatingPassword">Password</label>
            </div>

            <button className="w-100 btn btn-lg btn-primary" type="submit">
                Sign in
            </button>
            <p className="mt-5 mb-3 text-muted">&copy; 2021</p>
            </form>
        </div>
    )
}

export default Login
