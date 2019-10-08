import React from 'react';

function ClusterBar (props) {
    return (
        <nav className="navbar navbar-light bg-light">
            <a className="navbar-brand" href="/">{props.clusterName}</a>
        </nav>
    )
}

export default ClusterBar
