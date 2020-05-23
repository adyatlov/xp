import React from 'react';
import { BrowserRouter as Router, Switch, Route, useLocation} from "react-router-dom";

import TopBar from "./TopBar";
import ObjectSidePanel from "./ObjectSidePanel";
import Properties from "./Properties";
import ChildrenTable from "./ChildrenTable";
import Breadcrumb from "./Breadcrumb";
import {QueryRenderer} from "react-relay";
import environment from "./relayEnvironment";
import graphql from "babel-plugin-relay/macro";

export default function App() {
    return (
        <Router>
            <Switch>
                <Route path="/datasets/:datasetId/:objectId/:childrenTypeName">
                    <DatasetLayout mainPanel={<ChildrenTable/>}/>
                </Route>
                <Route path="/datasets/:datasetId/:objectId">
                    <DatasetLayout mainPanel={<Properties/>}/>
                </Route>
                <Route exact path="/">
                    <SelectDatasets/>
                </Route>
                <Route path="*">
                    <NoMatch/>
                </Route>
            </Switch>
        </Router>
    );
}

function DatasetLayout(props) {
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

function SelectDatasets() {
    return (
        <>
            <TopBar/>
            <div id="root" className="container">
                <div className="row mt-5">
                    <div className="col">
                        <div className="alert alert-light" role="alert">
                            <h4 className="alert-heading">Welcome to XP!</h4>
                            <p>Please, select a dataset. If there are no open datasets, or they are not what you need,
                                open a new one: insert a dataset URL, choose one of the compatible plugins,
                                and press "Open".</p>
                        </div>
                    </div>
                </div>
            </div>
        </>
    );
}

function NoMatch() {
    let location = useLocation();
    return (
        <>
            <TopBar/>
            <div id="root" className="container">
                <div className="row mt-5">
                    <div className="col">
                        <div className="alert alert-warning" role="alert">
                            <h4 className="alert-heading">Error</h4>
                            <p>Page <code>{location.pathname}</code> not found.</p>
                        </div>
                    </div>
                </div>
            </div>
        </>
    );
}
