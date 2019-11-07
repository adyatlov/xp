import React from "react"
import {Link} from "react-router-dom";

function ObjectLink(props) {
    let o = props.object;
    // onClick={props.handleSelectObject}
    return (
        <Link to={`/o/${o.type}/${o.id}`} href="/">{o.name}</Link>
    );
}

export default ObjectLink;