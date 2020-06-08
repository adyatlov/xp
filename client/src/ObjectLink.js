import React from 'react';
import {Link} from 'react-router-dom';

export default function ObjectLink(props) {
    const {objectId} = props;
    return(
        <Link to={"/o/" + objectId}>
            {props.children}
        </Link>
    );
}
