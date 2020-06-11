import React from 'react';
import {Link} from 'react-router-dom';

export default function ObjectLink(props) {
    const {nodeId, childTypeName, className} = props;
    if (!childTypeName) {
        return(
            <Link to={"/o/" + nodeId} className={className}>
                {props.children}
            </Link>
        );
    }
    return(
        <Link to={"/o/" + nodeId + "/" + childTypeName} className={className}>
            {props.children}
        </Link>
    );
}

