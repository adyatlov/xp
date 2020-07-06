import React from 'react';
import {useLocation, Route, Switch } from "react-router-dom";

import PageHome from "./PageHome";
import PageObject from "./PageObject";
import subscribeDatasetsChanged from "./subscribeDatasetsChanged";

export default class App extends React.Component {
    componentDidMount() {
        subscribeDatasetsChanged();
    }

    render() {
        return(
            <Switch>
                <Route path="/" exact>
                    <PageHome/>
                </Route>
                <Route path={["/o/:id/:childGroupTypeName", "/o/:id"]}>
                    <PageObject/>
                </Route>
                <Route>
                    <Error404 />
                </Route>
            </Switch>
        );
    }
}

function Error404() {
    const location = useLocation();
    return (
        <div className="alert alert-warning" role="alert">
            <h4 className="alert-heading">Error</h4>
            <p>Page <code>{location.pathname}</code> not found.</p>
        </div>
    );
}

