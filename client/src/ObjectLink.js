import React from "react"

function ObjectLink(props) {
    let o = props.object;
    return (
        <a data-otype={o.type} data-oid={o.id} href="/">{o.name}</a>
    );
}

export default ObjectLink;