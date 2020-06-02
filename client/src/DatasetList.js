import React from "react";
import {createFragmentContainer, commitMutation} from 'react-relay';
import graphql from 'babel-plugin-relay/macro';

function DatasetList(props) {
    if (!props.datasets) {
        return <div>Loading dataset list...</div>
    }
    return(
        <ul className="list-group">
            {props.datasets.map((value, index) => {
                return(
                    <li key={index} className="list-group-item">{value.root.type.name}: {value.root.name}</li>
                );
            })}
        </ul>
    );
}

export default createFragmentContainer(DatasetList, {
    datasets: graphql`
        fragment DatasetList_datasets on Dataset@relay(plural: true) {
            id
            root {
                name
                type {
                    name
                }
            }
        }
    `}
);
