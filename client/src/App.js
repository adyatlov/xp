import React from 'react';
import RootBar from './RootBar';
import graphql from 'babel-plugin-relay/macro';
import {QueryRenderer} from 'react-relay';
import {Environment, Network, RecordSource, Store} from 'relay-runtime';
import ObjectView from "./ObjectView";

function fetchQuery(
    operation,
    variables
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
        const params = this.props.match.params;
        let typeName = "";
        let objectId = "";
        if (params.typeName && params.objectId) {
            typeName = params.typeName;
            objectId = params.objectId;
        }
        return (
        <QueryRenderer
            environment={environment}
            query={graphql`
                     query App_Object_Query($typeName: String!, $objectId: String!) {
                        root {
                            ...RootBar_root
                        }
                        object(typeName: $typeName, objectId: $objectId) {
                            ...ObjectView_object
                        }
                     } 
                `}
            variables={{typeName: typeName, objectId: objectId}}
            render={({error, props}) => {
                if (error) {
                    return <div>Error!<br/>{error}</div>;
                }
                if (!props) {
                    return <div>Loading...</div>;
                }
                return (
                    <>
                        <RootBar root={props.root}/>
                        <ObjectView object={props.object}/>
                    </>
                );
            }}
        />);
    }
}

export default App;
