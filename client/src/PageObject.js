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
    query PageObjectQuery($nodeId: ID, $groupsIndexes: [Int!]) {
        node(id: $nodeId) {
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
                children {
                    id
                    index
                    type {
                        name
                        pluralName
                        description
                    }
                    totalCount
                }
                selectedGroups: children(indexes: $groupsIndexes) {
                    ...ChildrenPropertiesList_groups
                }
            }
        }
    }
`;

function PageObject(props) {
    const {object, groups} = props;
    if (!object) {
        return( <Error text="Unknown object"/> );
    }
    return (
        <div className="container-fluid">
            <div className="row">
                <div className="col-3">
                    <LeftPanel object={object}/>
                </div>
                <div className="col-9">
                    <ChildrenPropertiesList groups={groups} />
                </div>
            </div>
        </div>
    );
}

export default function PageObjectQuery() {
    let {nodeId, groupIndex} = useParams();
    let groupsIndexes = [0];
    if (groupIndex) {
        groupsIndexes = [parseInt(groupIndex)];
    }
    return(
        <QueryRenderer
            environment={environment}
            query={query}
            fetchPolicy={"store-and-network"}
            variables={{
                nodeId: nodeId,
                groupsIndexes: groupsIndexes,
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
                    <PageObject object={node} groups={node.selectedGroups} />
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
    const {children} = props;
    const params = useParams();
    const {nodeId} = params;
    return(
        <div className="list-group list-group-flush list-group-item-action">
            {children.map(group => (
                <ObjectLink nodeId={nodeId} groupIndex={group.index} key={group.id} className="list-group-item list-group-item-action d-flex justify-content-between align-items-center">
                    {group.type.pluralName}
                    <span className="badge badge-secondary badge-pill">{group.total}</span>
                </ObjectLink>))}
        </div>
    );
}


