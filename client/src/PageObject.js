import React from "react";
import graphql from 'babel-plugin-relay/macro';
import {createFragmentContainer} from 'react-relay';

function PageObject(props) {
    const {object} = props;
    if (!object) {
        return (<div>Object is loading...</div>);
    }
    return (
        <Layout>
            <div>{object.name}</div>
            <div>
                <ul>
                    {object.properties.map((property)=>{
                        return (
                            <li key={property.id}>{property.type.name}: {property.value}</li>
                        );
                    })}
                </ul>
            </div>
        </Layout>
    );
}
function Layout(props) {
    return(
        <div className="row">
            <div className="col-3">
                {props.children[0]}
            </div>
            <div className="col-9">
                {props.children[1]}
            </div>
        </div>
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


