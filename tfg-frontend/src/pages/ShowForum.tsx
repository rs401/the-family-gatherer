import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom';
import { Forum, Thread } from '../models/Models';

type ParamTypes = {
    id: string;
}
const ShowForum = () => {
    const [threads, setThreads] = useState<Thread[]>()
    const [forum, setForum] = useState<Forum>()
    const [loaded, setLoaded] = useState<boolean>(false)
    const params = useParams<ParamTypes>();
    let id = parseInt(params.id);

    const loadForum = async () => {
        const res = await fetch('http://127.0.0.1:8000/api/forum/'+id,{
            headers: {'Content-Type':'application/json'},
            credentials: 'include',
        });

        const data = await res.json();
        console.log(data.Threads)
        setForum(data);
        setThreads(data.Threads);
        setLoaded(true);
    }
    useEffect(() => {
        if(!loaded){
            loadForum();
        }
    });

    let theForum;
    if(loaded){
        theForum = (
            <div>
                <div className="forum-top">
                    { forum?.Name }
                </div>
                
                {threads?.map((thread, index) => {
                    return <div className="thread-list-item" key={index}>
                        a thread Id: { thread.ID } Title: {thread.Title} Body: {thread.Body}
                    </div>
                })}
                
            </div>
        );
    } else {
        theForum = (
            <h3>Loading...</h3>
        );
    }
    
    return (
        <div className="thread-list shadow rounded">
            { theForum }
        </div>
    )
}

export default ShowForum
