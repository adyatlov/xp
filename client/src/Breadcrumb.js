import React from 'react'

export default function Breadcrumb() {
    return(
        <nav aria-label="breadcrumb">
            <ol className="breadcrumb bg-white">
                <li className="breadcrumb-item"><a href="#">Production</a></li>
                <li className="breadcrumb-item"><a href="#">marathon</a></li>
                <li className="breadcrumb-item active" aria-current="page">accounting-service</li>
            </ol>
        </nav>
    );
}