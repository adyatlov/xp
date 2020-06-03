import graphql from "babel-plugin-relay/macro";
import {commitMutation} from "react-relay";
import environment from "./relayEnvironment";

const mutation = graphql`
    mutation removeDatasetMutation($id: ID!) {
        removeDataset(id: $id)
    }
`

export default function removeDataset(id) {
    return commitMutation(
        environment,
        {
            mutation,
            variables: {
                id: id,
            },
            onError: (error) => {
                console.error(error.message);
            },
        }
    )
}

