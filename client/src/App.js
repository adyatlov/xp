import React from 'react';
import {useLocation} from "react-router-dom";
import {QueryRenderer, requestSubscription} from 'react-relay';
import graphql from 'babel-plugin-relay/macro';

import environment from "./relayEnvironment";
import DatasetAdder from "./DatasetAdder";
import DatasetList from "./DatasetList";
// import Properties from "./Properties";
// import ChildrenTable from "./ChildrenTable";
// import {ObjectViewLayout} from "./ObjectViewLayout";

const query = graphql`
    query AppQuery {
        datasets {
            ...DatasetList_datasets
        }
    }
`;

const subscription = graphql`
    subscription AppSubscription {
        datasetsChanged {
            id
            root {
                name
                type {
                    name
                }
            }
        }
    }
`;

requestSubscription(
    environment,
    {
        subscription: subscription,
        variables: {},
        onCompleted: () => console.log("Server disconnected the subscription."),
        onError: error => console.error(error),
        updater: (store, data) => {
            let newDatasets = store.getPluralRootField("datasetsChanged")
            store.getRoot().setLinkedRecords(newDatasets, "datasets");
        },
    }
);

export default class App extends React.Component {
    render() {
        const {match} = this.props;
        return(
            <QueryRenderer
                environment={environment}
                query={query}
                // fetchPolicy={"store-and-network"}
                render={({error, props}) => {
                    if (error) {
                        return(
                            <Alert warning>
                                <h4 className="alert-heading">Error</h4>
                                <p>{error.text}</p>
                            </Alert>
                        );
                    }
                    if (!props) {
                        return (
                            <LoadingSpinner/>
                        );
                    }
                    let datasets = props.datasets;
                    if (match.path === "/") {
                        return (
                            <div id="root" className="container text-secondary">
                                <h4 className="mt-5 mb-4">Welcome to XP,
                                    <small> the explorer of heterogeneous datasets</small></h4>
                                <p>Please, open a new dataset:</p>
                                <DatasetAdder/>
                                {datasets.length > 0 &&  <p className="mt-3">or choose from the existing ones:</p>}
                                <DatasetList datasets={datasets}/>
                            </div>
                        );
                    }
                    return (
                        <>
                            <TopBar/>
                            <AlertNoMatch/>
                        </>
                    );
                }} />
    );
    }
}

function TopBar(props) {
    return (
        <nav className="navbar navbar-light bg-light">
            <form className="form-inline">
                {props.children}
            </form>
            <span className="navbar-brand">XP</span>
        </nav>
    );
}

function Alert(props) {
    let className = "alert alert-light";
    if (props.warning) {
        className = "alert alert-warning";
    }
    return(
        <div className={className} role="alert">
            {props.children}
        </div>
    );
}

function AlertNoMatch() {
    let location = useLocation();
    return (
        <Alert warning>
            <h4 className="alert-heading">Error</h4>
            <p>Page <code>{location.pathname}</code> not found.</p>
        </Alert>
    );
}

function LoadingSpinner() {
    return(
        <div style={{position: "fixed", top: "50%", left: "50%",
            transform:"translate(-50%, -50%)"}}>
            <h4 className="text-secondary">
                Loading XP...
            </h4>
            <div className="text-center">
                <div className="spinner-grow text-secondary mt-3" style={{width: "2rem", height: "2rem"}} role="status"/>
            </div>
        </div>
    );
}
