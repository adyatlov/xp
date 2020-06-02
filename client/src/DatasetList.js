import React from "react";
import {createFragmentContainer, commitMutation} from 'react-relay';
import graphql from 'babel-plugin-relay/macro';

function DatasetList(props) {
    if (!props.datasets) {
        return <div>Loading dataset list...</div>
    }
    return(
        <table className="table table-hover">
            <thead>
            <tr>
                <th scope="col">Type</th>
                <th scope="col">Name</th>
                <th scope="col">Plugin</th>
                <th scope="col">URL</th>
            </tr>
            </thead>
            <tbody>
            {props.datasets.map((value, index) => {
                return(
                    <tr key={index}>
                        <td>{value.root.type.name}</td>
                        <td>{value.root.name}</td>
                        <td>Plugin name goes here</td>
                        <td>http://url.goes.here/</td>
                    </tr>
                );
            })}
            </tbody>
        </table>
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
