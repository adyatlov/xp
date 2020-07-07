import {Environment, Network, RecordSource, Store} from "relay-runtime";
import {SubscriptionClient} from "subscriptions-transport-ws";
const hostPath = window.location.hostname+':7777/graphql';
function fetchQuery(
    operation,
    variables
) {
    // DEBUG:
    // console.log(operation.text);
    // console.log(variables);
    return fetch("http://"+hostPath, {
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
        // END DEBUG
        return response.json();
    });
}

function subscribeFunction(request, variables) {
    const query = request.text;
    const subscriptionClient = new SubscriptionClient('ws://'+hostPath, {reconnect: true});
    const client = subscriptionClient.request({query, variables});
    return client;
}

const network = Network.create(fetchQuery, subscribeFunction);
// const network = Network.create(fetchQuery);
const store = new Store(new RecordSource(), null);

const environment = new Environment({
    network,
    store
});

export default environment
