import {Environment, Network, RecordSource, Store} from "relay-runtime";

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
const store = new Store(new RecordSource(), null);

const environment = new Environment({
    network,
    store
});

export default environment
