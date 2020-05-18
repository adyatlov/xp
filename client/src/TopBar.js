import React from 'react'
import Plugin from "./Plugin";

export default function TopBar() {
    return (
        <nav className="navbar navbar-light bg-light">
            <form className="form-inline">
                <Plugin />
                <div className="btn-group ml-3">
                    <div className="btn-group">
                        <button className="btn btn-secondary dropdown-toggle" type="button" id="datasetsDropdown"
                                data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                            DC/OS Cluster Diagnostics Bundle: Production
                        </button>
                        <div className="dropdown-menu" aria-labelledby="datasetsDropdown">
                            <a className="dropdown-item" href="#">DC/OS Cluster Diagnostics Bundle: Development</a>
                            <a className="dropdown-item" href="#">DC/OS Cluster Diagnostics Bundle: Test</a>
                            <a className="dropdown-item" href="#">DC/OS Service Diagnostics Bundle: kafka-prod</a>
                        </div>
                    </div>
                    <button type="button" className="btn btn-dark text-nowrap">Close</button>
                </div>
            </form>
            <a className="navbar-brand">XP</a>
        </nav>
    );
}

