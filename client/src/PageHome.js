import React from "react";
import graphql from "babel-plugin-relay/macro";
import {QueryRenderer} from "react-relay";

import environment from "./relayEnvironment";
import DatasetAdder from "./DatasetAdder";
import DatasetList from "./DatasetList";
import PluginList from "./PluginList";
import LoadingSpinner from "./LoadingSpinner";
import Error from "./Error";

const query = graphql`
    query PageHomeQuery {
        allDatasets {
            ...DatasetList_datasets
        }
        allPlugins {
            ...PluginList_plugins
        }
    }
`;

function PageHome(props) {
    const {datasets, plugins} = props;
    return(
        <div id="root" className="container pt-4">
            <h1 className="mb-4">Welcome to XP
                <small className="text-muted"> &mdash; the explorer of heterogeneous datasets</small></h1>
            {props.children}
            <p className="text-secondary">Please open a new dataset:</p>
            <DatasetAdder/>
            {datasets.length > 0 &&
            <>
                <p className="mt-3 text-secondary">or choose an already loaded one:</p>
                <h2>Datasets</h2>
                <DatasetList datasets={datasets}/>
            </>}
            <p className="mt-3 text-secondary">To open a dataset of a particular type, you need a corresponding plugin.
                Below is the list of available plugins:</p>
            <h2>Plugins</h2>
            <PluginList plugins={plugins}/>
        </div>
    );
}

export default function PageHomeQuery() {
    return(
        <QueryRenderer
            environment={environment}
            query={query}
            render={({error, props}) => {
                if (error) {
                    console.error(error);
                    return(
                        <Error text={error.message} />
                    );
                }
                if (!props) {
                    return (
                        <LoadingSpinner />
                    );
                }
                const {allDatasets, allPlugins} = props;
                return(
                    <PageHome datasets={allDatasets} plugins={allPlugins} />
                );
            }}
        />
    );
}
