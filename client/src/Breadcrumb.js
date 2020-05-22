import React from 'react'

export default function Breadcrumb() {
    return(
        <nav aria-label="breadcrumb">
            <ol className="breadcrumb bg-white">
                <li className="breadcrumb-item">Production</li>
                <li className="breadcrumb-item">marathon</li>
                <li className="breadcrumb-item active" aria-current="page">accounting-service</li>
            </ol>
        </nav>
    );
}