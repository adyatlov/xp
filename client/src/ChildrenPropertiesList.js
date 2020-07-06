import React from "react";

import {useParams} from "react-router-dom";
import Error from "./Error";

export default function ChildrenPropertiesList(props) {
    const {childGroup} = props;
    const {childGroupTypeName} = useParams();
    if (!childGroup) {
        if(childGroupTypeName) {
            return(
                <Error text={'No child group of type "' + childGroupTypeName + '"'}/>
            )
        }
        return(
            <p>No child group selected</p>
        )
    }
    const propertyTypes = childGroup.type.propertyTypes;
    const edges = childGroup.children.edges;
    return(
        <table className="table">
            <thead>
            <tr>
                <th scope="col" className="text-nowrap">
                    {childGroup.type.name} Name
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

