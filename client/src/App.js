import React from 'react';
import {useLocation} from "react-router-dom";
import {QueryRenderer, requestSubscription} from 'react-relay';
import graphql from 'babel-plugin-relay/macro';

import environment from "./relayEnvironment";
import DatasetAdder from "./DatasetAdder";
import DatasetList from "./DatasetList";
import PluginList from "./PluginList";
// import Properties from "./Properties";
// import ChildrenTable from "./ChildrenTable";
// import {ObjectViewLayout} from "./ObjectViewLayout";

const query = graphql`
    query AppQuery {
        datasets {
            ...DatasetList_datasets
        }
        plugins {
            ...PluginList_plugins
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
        updater: (store) => {
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
                    let plugins = props.plugins;
                    if (match.path === "/" && match.isExact) {
                        return (
                            <PageHome datasets={datasets} plugins={plugins}>
                            </PageHome>
                        );
                    }
                    return (
                        <PageHome datasets={datasets} plugins={plugins}>
                            <Page404/>
                        </PageHome>
                    );
                }}
            />
        );
    }
}

// function TopBar(props) {
//     return (
//         <nav className="navbar navbar-light bg-light">
//             <form className="form-inline">
//                 {props.children}
//             </form>
//             <span className="navbar-brand">XP</span>
//         </nav>
//     );
// }

function PageHome(props) {
    const {datasets} = props;
    return(
        <div id="root" className="container pt-4">
            <h1 className="mb-4">Welcome to XP
                <small className="text-muted"> &mdash; the explorer of heterogeneous datasets</small></h1>
            {props.children}
            <p className="text-secondary">Please, open a new dataset:</p>
            <DatasetAdder/>
            {datasets.length > 0 &&
            <>
                <p className="mt-3 text-secondary">or choose the already loaded one:</p>
                <h2>Datasets</h2>
                <DatasetList datasets={datasets}/>
            </>}
            <h2>Plugins</h2>
            <PluginList plugins={props.plugins}/>
        </div>
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