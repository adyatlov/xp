import React from 'react';
import ClusterBar from './ClusterBar';
import graphql from 'babel-plugin-relay/macro';
import {QueryRenderer} from 'react-relay';
import {Environment, Network, RecordSource, Store} from 'relay-runtime';

function fetchQuery(
    operation,
    variables,
    cacheConfig,
    uploadables,
) {
    return fetch('http://localhost:7777/graphql', {
        method: 'POST',
        headers: {
            // Add authentication and other headers here
            'content-type': 'application/json'
        },
        body: JSON.stringify({
            query: operation.text, // GraphQL text from input
            variables,
        }),
    }).then(response => {
        return response.json();
    });
}

const network = Network.create(fetchQuery);
const store = new Store(new RecordSource());

const environment = new Environment({
    network,
    store
});

class App extends React.Component {
    render() {
        return (
        <QueryRenderer
            environment={environment}
            query={graphql`
                     query App_RootName_Query {
                        root {
                            name
                        }
                     } 
                `}
            variables={{}}
            render={({error, props}) => {
                if (error) {
                    return <div>Error!<br/>{error}</div>;
                }
                if (!props) {
                    return <div>Loading...</div>;
                }
                return (
                    <>
                        <ClusterBar clusterName={props.root.name}/>
                        {/*<ObjectView/>*/}
                    </>
                );
            }}
        />);
    }
}

export default App;
