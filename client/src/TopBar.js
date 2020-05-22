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
                            <button className="dropdown-item" id={1}>DC/OS Cluster Diagnostics Bundle: Development</button>
                            <button className="dropdown-item" id={2}>DC/OS Cluster Diagnostics Bundle: Test</button>
                            <button className="dropdown-item" id={3}>DC/OS Service Diagnostics Bundle: kafka-prod</button>
                        </div>
                    </div>
                    <button type="button" className="btn btn-dark text-nowrap">Close</button>
                </div>
            </form>
            <span className="navbar-brand">XP</span>
        </nav>
    );
}

