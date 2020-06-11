import graphql from "babel-plugin-relay/macro";
import {requestSubscription} from "react-relay";
import environment from "./relayEnvironment";

const subscription = graphql`
    subscription subscribeDatasetsChangedSubscription {
        datasetUpdated {
            eventType
            idToRemove
            dataset {
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
                    children {
                        type {
                            name
                        }
                    }
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
                const datasetUpdated = store.getRoot().getLinkedRecord("datasetUpdated");
                const eventType = datasetUpdated.getValue("eventType") ;
                let datasetRecords = store.getRoot().getLinkedRecords("allDatasets");
                switch (eventType) {
                    case "added":
                        const datasetRecord = datasetUpdated.getLinkedRecord("dataset");
                        datasetRecords.push(datasetRecord);
                        break;
                    case "removed":
                        const idToRemove = datasetUpdated.getValue("idToRemove");
                        datasetRecords = datasetRecords.filter(datasetRecord => {
                            return datasetRecord.getValue("id") !== idToRemove;
                        })
                        break;
                    default:
                        console.warn("Unknown eventType: ", eventType);
                        return
                }
                store.getRoot().setLinkedRecords(datasetRecords, "allDatasets")
            },
        }
    );
}
