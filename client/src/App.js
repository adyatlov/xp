import React from 'react';
import {useLocation} from "react-router-dom";
import {QueryRenderer} from 'react-relay';
import graphql from 'babel-plugin-relay/macro';

import environment from "./relayEnvironment";
import subscribeDatasetsChanged from "./subscribeDatasetsChanged";

import Error from "./Error";
import LoadingSpinner from "./LoadingSpinner";
import PageHome from "./PageHome";
import PageObject from "./PageObject";

const query = graphql`
    query AppQuery {
        allDatasets {
            ...DatasetList_datasets
        }
        allPlugins {
            ...PluginList_plugins
        }
    }
`;

export default class App extends React.Component {
    componentDidMount() {
        subscribeDatasetsChanged();
    }

    render() {
        const {match} = this.props;
        return(
            <QueryRenderer
                environment={environment}
                query={query}
                fetchPolicy={"store-and-network"}
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
                    let datasets = props.allDatasets;
                    let plugins = props.allPlugins;
                    if (match.path === "/" && match.isExact) {
                        return (
                            <PageHome datasets={datasets} plugins={plugins} />
                        );
                    }
                    if (match.path === "/o/:datasetId/:objectId" && match.isExact) {
                        return (
                            <PageObject />
                        )
                    }
                    return (
                        <PageHome datasets={datasets} plugins={plugins}>
                            <Error404 />
                        </PageHome>
                    );
                }}
            />
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

