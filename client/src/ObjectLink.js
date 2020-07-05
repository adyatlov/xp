import React from 'react';
import {Link} from 'react-router-dom';

export default function ObjectLink(props) {
    const {nodeId, groupIndex, className} = props;
    if (!groupIndex) {
        return(
            <Link to={"/o/" + nodeId} className={className}>
                {props.children}
            </Link>
        );
    }
    return(
        <Link to={"/o/" + nodeId + "/" + groupIndex} className={className}>
            {props.children}
        </Link>
    );
}
