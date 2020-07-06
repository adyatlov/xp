import React from "react";
import {Link} from 'react-router-dom';

export default function TopBar(props) {
    return (
        <nav className="navbar navbar-light bg-light mb-3">
            <Link to="/" className="navbar-brand">XP</Link>
        </nav>
    );
}

