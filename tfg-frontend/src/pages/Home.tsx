import React, { useEffect } from 'react'

const Home = (props: {name: string, setName: (name: string) => void}) => {
    useEffect(() => {
        (
            async () => {
                const res = await fetch('http://127.0.0.1:8000/api/user',{
                    headers: {'Content-Type':'application/json'},
                    credentials: 'include',
                });

                const data = await res.json();

                props.setName(data.name)
            }
        )();
    });
    return (
        <div>
            {props.name ? 'Hello, ' + props.name : 'Hello, Guest' }
        </div>
    )
}
export default Home