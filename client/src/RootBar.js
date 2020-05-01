import React from 'react';

import {createFragmentContainer} from 'react-relay';
import graphql from 'babel-plugin-relay/macro';
import ObjectLink from "./ObjectLink";


class RootBar extends React.Component {
    render() {
        const {typeName, objectId, name} = this.props.root;
        return (
            <nav className="navbar navbar-light bg-light">
                <ObjectLink text={typeName + " " + name}
                            typeName={typeName}
                            objectId={objectId}
                            className="navbar-brand"
                />
            </nav>
        )
    }
}

export default createFragmentContainer(
    RootBar,
    {
        root: graphql`
            fragment  RootBar_root on Object {
                typeName
                objectId
                name
            }
        `
    },
)

