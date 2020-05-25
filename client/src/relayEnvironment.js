import {Environment, Network, RecordSource, Store} from "relay-runtime";

function fetchQuery(
    operation,
    variables
) {
    // DEBUG:
    // console.log(operation.text);
    // console.log(variables);
    // let debug = JSON.stringify({
    //     query: operation.text, // GraphQL text from input
    //     variables,
    // });
    // console.log(debug);
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
        // DEBUG:
        // let tmp = response.json();
        // tmp.then((val) => console.log(val.data));
        // return tmp;
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
