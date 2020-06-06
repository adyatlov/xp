import React from 'react';
import {Link} from 'react-router-dom';

export default function ObjectLink(props) {
    const {datasetId, objectId} = props;
    return(
        <Link to={"/o/" + datasetId + "/" + objectId}>
            {props.children}
        </Link>
    );
}
