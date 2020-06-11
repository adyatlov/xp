import React from "react";
import graphql from 'babel-plugin-relay/macro';
import {QueryRenderer} from 'react-relay';
import {useParams} from "react-router-dom";

import environment from "./relayEnvironment";
import LoadingSpinner from "./LoadingSpinner";
import Error from "./Error";
import ChildrenPropertiesList from "./ChildrenPropertiesList";
import ObjectLink from "./ObjectLink";

const query = graphql`
    query PageObjectQuery($nodeId: ID, $childTypeNames: [String!]) {
        node(id: $nodeId) {
            id
            ... on Object {
                name
                type {
                    name
                    description
                }
                properties {
                    id
                    value
                    type {
                        name
                        description
                    }
                }
                children {
                    id
                    type {
                        name
                        pluralName
                        description
                    }
                    total
                }
                selectedChildren: children(typeNames: $childTypeNames) {
                    ...ChildrenPropertiesList_childrenProperties
                }
            }
        }
    }
`;

function PageObject(props) {
    const {object} = props;
    const {selectedChildren} = object;
    return (
        <div className="container-fluid">
            <div className="row">
                <div className="col-3">
                    <LeftPanel object={object}/>
                </div>
                <div className="col-9">
                    <ChildrenPropertiesList childrenProperties={selectedChildren} />
                </div>
            </div>
        </div>
    );
}

export default function PageObjectQuery() {
    const params = useParams();
    let {nodeId, childTypeNames} = params;
    if (childTypeNames) {
        childTypeNames = [childTypeNames];
    }
    return(
        <QueryRenderer
            environment={environment}
            query={query}
            fetchPolicy={"store-and-network"}
            variables={{
                nodeId: nodeId,
                childTypeNames: childTypeNames,
            }}
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
                const {node, childrenProperties} = props;
                return(
                    <PageObject object={node} childrenProperties={childrenProperties} />
                );
            }}
        />
    );
}

function LeftPanel(props) {
    const {object} = props;
    const {properties, children} = object;
    return (
        <div className="card">
            <div className="card-header">
                <h5 className="align-middle">{object.name} <span className="badge badge-primary">{object.type.name}</span></h5>
            </div>
            <div className="card-body">
                <PropertyList properties={properties} />
            </div>
            <GroupList children={children}/>
        </div>
    );
}

function PropertyList(props) {
    const {properties} = props;
    return(
        <table className="table table-sm table-borderless">
            <tbody>
            {properties.map(property => (
                <tr key={property.id}>
                    <th>{property.type.name} </th>
                    <td>{property.value}</td>
                </tr>
            ))}
            </tbody>
        </table>
    );
}

function GroupList(props) {
    const {children} = props;
    const params = useParams();
    const {nodeId} = params;
    return(
        <div className="list-group list-group-flush list-group-item-action">
            {children.map(group => (
                <ObjectLink nodeId={nodeId} childTypeName={group.type.name} key={group.id} className="list-group-item list-group-item-action d-flex justify-content-between align-items-center">
                    {group.type.pluralName}
                    <span className="badge badge-secondary badge-pill">{group.total}</span>
                </ObjectLink>))}
        </div>
    );
}


