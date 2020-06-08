import React from "react";
import graphql from 'babel-plugin-relay/macro';
import {createFragmentContainer} from 'react-relay';

function PageObject(props) {
    const {object} = props;
    if (!object) {
        return (<div>Object is loading...</div>);
    }
    return (
        <div className="container-fluid">
            <div className="row">
                <div className="col-3">
                    <LeftPanel object={object}/>
                </div>
                <div className="col-9">
                    Children table goes here
                </div>
            </div>
        </div>
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
            <ChildrenGroupList children={children}/>
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

function ChildrenGroupList(props) {
    const {children} = props;
    return(
        <ul className="list-group list-group-flush">
            {children.map(group => (
            <li key={group.id} className="list-group-item d-flex justify-content-between align-items-center">
                {group.type.pluralName}
                <span className="badge badge-secondary badge-pill">{group.total}</span>
            </li>))}
        </ul>
    );
}

export default createFragmentContainer(PageObject, {
    object: graphql`
        fragment PageObject_object on Node {
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
                        valueType
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
            }
        }
    `
});


