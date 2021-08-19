import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import {Forum} from '../models/Models';

const Forums = () => {
    const [forums, setForums] = useState<Forum[]>([])
    const [loaded, setLoaded] = useState<boolean>(false)

    const getForums = async () => {
        const res = await fetch('http://127.0.0.1:8000/api/forum',{
            headers: {'Content-Type':'application/json'},
            credentials: 'include',
        });

        const data = await res.json();
        console.log(data)

        setForums(data)
        setLoaded(true)
    };
    useEffect(() => {
        if(!loaded){
            getForums();
        }
    });

    let theStuff;
    if(loaded){
        theStuff = (
            <div >
            { forums.map((forum, index) => {
                return <Link to={`/forum/${forum.ID}`} key={index}>
                    <div className="forum-list-item " key={index}>
                        { forum.Name } - fid: { forum.ID } - uid: {forum.UserId}
                    </div>
                    </Link>;
            }) }
            </div>
        )

    } else {
        <h3>Loading...</h3>
    }

    return (
        <div className="forum-list shadow rounded ">
            <div className="row justify-content-between">
                <div className="col-2">
                    Forums
                </div>
                <div className="col-2 text-end">
                    New Forum
                </div>
            </div>
            
            { theStuff }
            
        </div>
    )
}

export default Forums
