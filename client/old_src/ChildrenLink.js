import React from "react"
import {Link} from "react-router-dom";

function ChildrenLink(props) {
    const t = props.typeName;
    const id = props.objectId;
    const ct = props.childrenTypeName;
    return (
        <Link to={`/o/${t}/${id}/${ct}`}>{props.children}</Link>
    );
}

export default ChildrenLink;
