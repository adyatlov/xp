import React from 'react';
import { BrowserRouter as Router, Switch, Route, useLocation} from "react-router-dom";

import TopBar from "./TopBar";
import ObjectSidePanel from "./ObjectSidePanel";
import Properties from "./Properties";
import ChildrenTable from "./ChildrenTable";
import Breadcrumb from "./Breadcrumb";

export default function App() {
    return (
        <Router>
            <Switch>
                <Route path="/properties/:datasetId/:objectId">
                    <Layout mainPanel={<Properties/>} />
                </Route>
                <Route path="/children/:datasetId/:objectId/:childrenTypeName">
                    <Layout mainPanel={<ChildrenTable />} />
                </Route>
                <Route exact path="/">
                    <Layout message={<NoDatasets />}/>
                </Route>
                <Route path="*">
                    <Layout message={<NoMatch />} />
                </Route>
            </Switch>
        </Router>
    );
}

function Layout(props) {
    if (props.mainPanel) {
        return (
            <>
                <TopBar/>
                <div id="root" className="container-fluid">
                    <Breadcrumb/>
                    <div className="row">
                        <div className="col-3"><ObjectSidePanel /></div>
                        <div className="col-9">{props.mainPanel}</div>
                    </div>
                </div>
            </>
        );
    }
    return (
        <>
            <TopBar/>
            <div id="root" className="container">
                {props.message}
            </div>
        </>
    );
}

function NoDatasets() {
    return (
        <div className="row mt-5">
            <div className="col">
                <div className="alert alert-light" role="alert">
                    <h4 className="alert-heading">Welcome to XP!</h4>
                    <p>There are no open datasets at the moment. Please insert a dataset URL, choose one of the compatible
                        plugins, and press "Open".</p>
                </div>
            </div>
        </div>
    );
}

function NoMatch() {
    let location = useLocation();
    return (
        <div className="row mt-5">
            <div className="col">
                <div className="alert alert-warning" role="alert">
                    <h4 className="alert-heading">Error</h4>
                    <p>Page <code>{location.pathname}</code> not found.</p>
                </div>
            </div>
        </div>
    );
}
