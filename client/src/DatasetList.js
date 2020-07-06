import React from "react";
import {createFragmentContainer} from 'react-relay';
import graphql from 'babel-plugin-relay/macro';

import removeDataset from "./removeDataset";
import ObjectLink from "./ObjectLink";

function DatasetList(props) {
    if (!props.datasets) {
        return <div>Loading dataset list...</div>
    }
    const {datasets} = props;
    console.log(datasets);
    return(
        <table className="table table-hover">
            <thead>
            <tr>
                <th scope="col">Name</th>
                <th scope="col">Type</th>
                <th scope="col">Plugin</th>
                <th scope="col">URL</th>
                <th scope="col">Added</th>
                <th scope="col"/>
            </tr>
            </thead>
            <tbody>
            {datasets.map((dataset) => {
                let added = parseInt(dataset.added) * 1000;
                const firstAvailableChildGroupTypeName = dataset.root.firstAvailableChildGroupTypeName;
                added = new Date(added);
                added = added.toLocaleString();
                return(
                    <tr key={dataset.id}>
                        <td className="text-nowrap">
                            <ObjectLink id={dataset.root.id}
                                        childGroupTypeName={firstAvailableChildGroupTypeName}>
                                {dataset.root.name}
                            </ObjectLink>
                        </td>
                        <td>{dataset.root.type.name}</td>
                        <td>{dataset.plugin.name}</td>
                        <td>{dataset.url}</td>
                        <td className="text-nowrap">{added}</td>
                        <td className="align-middle">
                            <RemoveDatasetButton id={dataset.id}/>
                        </td>
                    </tr>
                );
            })}
            </tbody>
        </table>
    );
}

function RemoveDatasetButton(props) {
    const onClick = () => {
        removeDataset(props.id);
    }
    return (
        <button onClick={onClick}
                type="button"
                className="close text-white dataset-list-close-button"
                aria-label="Close">
            <span aria-hidden="true">&times;</span>
        </button>
    );
}
export default createFragmentContainer(DatasetList, {
    datasets: graphql`
        fragment DatasetList_datasets on Dataset@relay(plural: true) {
            id
            plugin {
                name
            }
            url
            added
            root {
                id
                name
                type {
                    name
                }
                firstAvailableChildGroupTypeName 
            }
        }
    `}
);
