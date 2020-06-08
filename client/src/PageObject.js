import React from "react";
import graphql from 'babel-plugin-relay/macro';
import {QueryRenderer} from 'react-relay';
import {useParams} from "react-router-dom";

import environment from "./relayEnvironment";
import Error from "./Error";
import LoadingSpinner from "./LoadingSpinner";

const query = graphql`
    query PageObjectQuery($objectId: ID!) {
        node(id: $objectId) {
            id
            ... on Object {
                name
                type {
                    name
                    description
                }
                properties {
                    value
                    type {
                        name
                        valueType
                        description
                    }
                }
                children {
                    type {
                        name
                        pluralName
                        description
                    }
                    total
                }
            }
        }
    }`;

export default function PageObjectQuery(props) {
    const {datasetId, objectId} = useParams();
    return (
        <QueryRenderer
            environment={environment}
            query={query}
            fetchPolicy={"store-and-network"}
            variables={{datasetId: datasetId, objectId: objectId}}
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
                return (
                    <PageObject object={props.object}/>
                );
            }}
        />
    );
}

function PageObject(props) {
    return (
        <div>{props.object.name}</div>
    );
}

