import graphql from "babel-plugin-relay/macro";
import {requestSubscription} from "react-relay";
import environment from "./relayEnvironment";

const subscription = graphql`
    subscription subscribeDatasetsChangedSubscription {
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

export default function subscribeDatasetsChanged() {
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
}
