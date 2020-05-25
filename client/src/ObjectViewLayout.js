import React from "react";
import {QueryRenderer} from 'react-relay';
import graphql from 'babel-plugin-relay/macro';
import environment from "./relayEnvironment";

import Breadcrumb from "./Breadcrumb";
import ObjectSidePanel from "./ObjectSidePanel";

const query = graphql`
    query ObjectViewLayoutQuery($datasetId: ID!, $objectId: ID!) {
        object(datasetId: $datasetId, id: $objectId) {
            id
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
    }`;

export function ObjectViewLayout(props) {
    return (
            <QueryRenderer
                environment={environment}
                query={query}
                variables={{}}
                render={({error, props}) => {
                    return(
                        <div id="root" className="container-fluid">
                            <Breadcrumb/>
                            <div className="row">
                                <div className="col-3"><ObjectSidePanel/></div>
                                <div className="col-9">props.children</div>
                            </div>
                        </div>
                    );
                }}/>
    )
}