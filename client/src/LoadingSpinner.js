import React from 'react'

export default function LoadingSpinner(props) {
    let {message} = props
    if (!message) {
        message = "Loading..."
    }
    return(
        <div style={{position: "fixed", top: "50%", left: "50%",
            transform:"translate(-50%, -50%)"}}>
            <h4 className="text-secondary">
                {message}
            </h4>
            <div className="text-center">
                <div className="spinner-grow text-secondary mt-3" style={{width: "2rem", height: "2rem"}} role="status"/>
            </div>
        </div>
    );
}
