import React from "react";
import graphql from 'babel-plugin-relay/macro';
import {QueryRenderer} from 'react-relay';
import {useParams} from "react-router-dom";

import environment from "./relayEnvironment";
import TopBar from "./TopBar";
import LoadingSpinner from "./LoadingSpinner";
import Error from "./Error";
import ChildrenPropertiesList from "./ChildrenPropertiesList";
import ObjectLink from "./ObjectLink";

const query = graphql`
    query PageObjectQuery($id: ID, $childGroupTypeName: String) {
        node(id: $id) {
            id
            ... on Object {
                name
                type {
                    name
                    description
                }
                properties {
                    edges {
                        node {
                            id
                            value
                            type {
                                name
                                description
                            }
                        }
                    }
                }
                childGroups {
                    id
                    type {
                        name
                        pluralName
                        description
                    }
                    totalCount
                }
                childGroup(typeName: $childGroupTypeName) {
                    type {
                        name
                        propertyTypes {
                            name
                        }
                    }
                    children {
                        edges {
                            node {
                                id
                                name
                                properties {
                                    edges {
                                        node {
                                            id
                                            value
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    }
`;

function PageObject(props) {
    const {object} = props;
    if (!object) {
        return( <Error text="Unknown object"/> );
    }
    return (
        <div className="container-fluid">
            <TopBar/>
            <div className="row">
                <div className="col-3">
                    <LeftPanel object={object}/>
                </div>
                <div className="col-9">
                    <ChildrenPropertiesList childGroup={object.childGroup} />
                </div>
            </div>
        </div>
    );
}

export default function PageObjectQuery() {
    let {id, childGroupTypeName} = useParams();
    return(
        <QueryRenderer
            environment={environment}
            query={query}
            fetchPolicy={"store-and-network"}
            variables={{
                id: id,
                childGroupTypeName: childGroupTypeName,
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
                console.log(props);
                const {node} = props;
                return(
                    <PageObject object={node}/>
                );
            }}
        />
    );
}

function LeftPanel(props) {
    const {object} = props;
    console.log(object)
    const {properties, childGroups} = object;
    return (
        <div className="card">
            <div className="card-header">
                <h5 className="align-middle">{object.name} <span className="badge badge-primary">{object.type.name}</span></h5>
            </div>
            <div className="card-body">
                <PropertyList properties={properties} />
            </div>
            <GroupList childGroups={childGroups}/>
        </div>
    );
}

function PropertyList(props) {
    const {properties} = props;
    return(
        <table className="table table-sm table-borderless">
            <tbody>
            {properties.edges.map(edge => (
                <tr key={edge.node.id}>
                    <th>{edge.node.type.name} </th>
                    <td>{edge.node.value}</td>
                </tr>
            ))}
            </tbody>
        </table>
    );
}

function GroupList(props) {
    const {childGroups} = props;
    const params = useParams();
    const {id} = params;
    return(
        <div className="list-group list-group-flush list-group-item-action">
            {childGroups.map(group => (
                <ObjectLink id={id}
                            childGroupTypeName={group.type.name}
                            key={group.id}
                            className="list-group-item list-group-item-action d-flex justify-content-between align-items-center">
                    {group.type.pluralName}
                    <span className="badge badge-primary badge-pill">{group.totalCount}</span>
                </ObjectLink>))}
        </div>
    );
}
