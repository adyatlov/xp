import React from "react";
import {createFragmentContainer} from 'react-relay';
import graphql from 'babel-plugin-relay/macro';

function PluginList(props) {
    const plugins = props.plugins;
    if (!plugins) {
        return <div>Loading plugin list...</div>
    }
    return(
        <table className="table">
            <thead>
            <tr>
                <th scope="col">Name</th>
                <th scope="col">Description</th>
            </tr>
            </thead>
            <tbody>
            {plugins.map((plugin) => {
                return(
                    <tr key={plugin.name}>
                        <td className="text-nowrap">{plugin.name}</td>
                        <td>{plugin.description}</td>
                    </tr>
                );
            })}
            </tbody>
        </table>
    );
}

export default createFragmentContainer(PluginList, {
    plugins: graphql`
        fragment PluginList_plugins on Plugin@relay(plural: true) {
            id
            name
            description
        }
    `}
);
