import React from "react";
import graphql from 'babel-plugin-relay/macro';
import {createFragmentContainer} from 'react-relay';

function PageObject(props) {
    const {object} = props;
    if (!object) {
        return (<div>Object is loading...</div>);
    }
    return (
        <div>{props.object.name}</div>
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
    `
});

