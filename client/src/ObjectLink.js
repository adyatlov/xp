import React from 'react';
import {Link} from 'react-router-dom';

export default function ObjectLink(props) {
    const {id, childGroupTypeName, className} = props;
    if (!childGroupTypeName) {
        return(
            <Link to={"/o/" + id} className={className}>
                {props.children}
            </Link>
        );
    }
    return(
        <Link to={"/o/" + id + "/" + childGroupTypeName}
              className={className}>
            {props.children}
        </Link>
    );
}
