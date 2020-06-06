import React from "react";

export default function Error(props) {
    return(
        <div className="alert alert-warning" role="alert">
            <h4 className="alert-heading">Error</h4>
            <p>{props.text}</p>
        </div>
    );
}
