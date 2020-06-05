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
            plugin {
                name
            }
            url
            added
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
                fetchPolicy={"store-and-network"}
                render={({error, props}) => {
                    if (error) {
                        return(
                            <Page500 text={error.text} />
                        );
                    }
                    if (!props) {
                        return (
                            <LoadingSpinner/>
                        );
                    }
                    let datasets = props.datasets;
                    if (match.path === "/" && match.isExact) {
                        return (
                            <PageHome datasets={datasets}>
                            </PageHome>
                        );
                    }
                    return (
                        <PageHome datasets={datasets}>
                            <Page404/>
                        </PageHome>
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

function Page404() {
    let location = useLocation();
    return (
        <div className="alert alert-warning" role="alert">
            <h4 className="alert-heading">Error</h4>
            <p>Page <code>{location.pathname}</code> not found.</p>
        </div>
    );
}

function Page500(props) {
    return(
        <div className="alert alert-warning" role="alert">
            <h4 className="alert-heading">Error</h4>
            <p>{props.text}</p>
        </div>
    );
}

function PageHome(props) {
    const {datasets} = props;
    return(
        <div id="root" className="container text-secondary pt-5">
            <h4 className="mb-4">Welcome to XP,
                <small> the explorer of heterogeneous datasets</small></h4>
            {props.children}
            <p>Please, open a new dataset:</p>
            <DatasetAdder/>
            {datasets.length > 0 &&
            <>
                <p className="mt-3">or choose the existing one:</p>
                <DatasetList datasets={datasets}/>
            </>}
        </div>
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
