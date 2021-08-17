import React, { SyntheticEvent, useState } from 'react'
import { Redirect } from 'react-router-dom';

const Register = () => {
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [password1, setPass1] = useState('');
    const [password2, setPass2] = useState('');
    const [redirect, setRedirect] = useState(false);

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();
        await fetch('http://127.0.0.1:8000/api/register',{
            method: 'POST',
            headers: {'Content-Type':'application/json'},
            body: JSON.stringify({
                name,
                email,
                password1,
                password2,
            })
        });
        setRedirect(true);
    }

    if(redirect){
        return <Redirect to="/login" />
    }

    return (
        <div className="form-signin">
            <form onSubmit={submit}>
            <h1 className="h3 mb-3 fw-normal">Please register</h1>

            <div className="form-floating">
                <input
                type="text"
                className="form-control"
                id="floatingInput0"
                name="name"
                placeholder="Display Name"
                onChange={e => setName(e.target.value)}
                />
            <label htmlFor="floatingInput">Display Name</label>
            </div>
            <span className="small">This is what people will see.</span>
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
                id="floatingPassword1"
                name="password1"
                placeholder="Password"
                onChange={e => setPass1(e.target.value)}
                />
            <label htmlFor="floatingPassword1">Password</label>
            </div>
            <div className="form-floating">
                <input
                type="password"
                className="form-control"
                id="floatingPassword2"
                name="password2"
                placeholder="Password"
                onChange={e => setPass2(e.target.value)}
                />
            <label htmlFor="floatingPassword2">Password (again)</label>
            </div>

            <button className="w-100 btn btn-lg btn-primary" type="submit">
                Submit
            </button>
            <p className="mt-5 mb-3 text-muted">&copy; 2021</p>
            </form>
        </div>
    )
}
export  default Register