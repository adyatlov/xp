import React from "react"
import {Link} from "react-router-dom";

function ObjectLink(props) {
    const t = props.typeName;
    const id = props.objectId;
    return (
        <Link to={`/o/${t}/${id}`} >{props.children}</Link>
    );
}

export default ObjectLink;