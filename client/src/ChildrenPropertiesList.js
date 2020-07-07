import React from "react";

import {useParams} from "react-router-dom";
import Error from "./Error";
import ObjectLink from "./ObjectLink";
import PropertyValue from "./PropertyValue";

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
                <th scope="col" className="text-nowrap text-capitalize">
                    {childGroup.type.name} Name
                </th>
                {propertyTypes.map(propertyType => {
                    return(
                        <th key={propertyType.name} scope="col" className="text-capitalize">
                            {propertyType.name}
                        </th>
                    );
                })}
            </tr>
            </thead>
            <tbody>
            {edges.map((edge) => {
                const object = edge.node
                return(
                    <tr key={object.id}>
                        <td>
                            <ObjectLink id={object.id} childGroupTypeName={object.firstAvailableChildGroupTypeName}>
                            {object.name}
                            </ObjectLink>
                        </td>
                        {object.properties.edges.map((edge, index) => {
                            const property = edge.node
                            return(
                                <td key={property.id}>
                                    <PropertyValue value={property.value}
                                                   type={propertyTypes[index].valueType} />
                                </td>
                            );
                        })}
                    </tr>
                );
            })}
            </tbody>
        </table>
    );
}

