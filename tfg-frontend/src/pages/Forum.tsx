import React from 'react'
import { useParams } from 'react-router-dom';

type ParamTypes = {
    id: string;
}
const Forum = () => {
    const params = useParams<ParamTypes>();
    let id = parseInt(params.id);
    
    return (
        <div>
            a forum { id }
        </div>
    )
}

export default Forum
