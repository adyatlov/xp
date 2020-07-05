import {createFragmentContainer} from "react-relay";
import graphql from "babel-plugin-relay/macro";
import React from "react";

import {useParams} from "react-router-dom";
import Error from "./Error";

function ChildrenPropertiesList(props) {
    let {groups} = props;
    const {groupIndex} = useParams();
    if (!groups) {
        if(!groupIndex)
        return(
            <Error text={'No children group with the index "' + groupIndex + '"'}/>
        )
        return
    }
    let group = groups[0]
    const propertyTypes = group.type.properties;
    console.log(group)
    const edges = group.objects.edges;
    return(
        <table className="table">
            <thead>
            <tr>
                <th scope="col" className="text-nowrap">
                    {group.type.name} name
                </th>
                {propertyTypes.map(propertyType => {
                    return(
                        <th key={propertyType.name} scope="col">{propertyType.name}</th>
                    );
                })}
            </tr>
            </thead>
            <tbody>
            {edges.map((edge) => {
                const object = edge.node
                return(
                    <tr key={object.id}>
                        <td>{object.name}</td>
                        {object.properties.edges.map((edge) => {
                            const property = edge.node
                            return(
                                <td key={property.id}>{property.value}</td>
                            );
                        })}
                    </tr>
                );
            })}
            </tbody>
        </table>
    );
}

export default createFragmentContainer(ChildrenPropertiesList, {
    groups: graphql`
        fragment ChildrenPropertiesList_groups on ObjectGroup@relay(plural: true) {
            type {
                name
                properties {
                    name
                }
            }
            objects {
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
    `
});

