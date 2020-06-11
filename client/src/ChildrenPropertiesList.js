import {createFragmentContainer} from "react-relay";
import graphql from "babel-plugin-relay/macro";
import React from "react";

function ChildrenPropertiesList(props) {
    let {childrenProperties} = props;
    childrenProperties = childrenProperties[0]
    const childrenType = childrenProperties.type;
    const propertyTypes = childrenType.properties;
    const objects = childrenProperties.objects;
    return(
        <table className="table">
            <thead>
            <tr>
                <th scope="col" className="text-nowrap">
                    {childrenProperties.type.name} name
                </th>
                {propertyTypes.map((propertyType) => {
                    return(
                        <th key={propertyType.name} scope="col">{propertyType.name}</th>
                    );
                })}
            </tr>
            </thead>
            <tbody>
            {objects.map((object) => {
                return(
                    <tr key={object.id}>
                        <td>{object.name}</td>
                        {object.properties.map((property) => {
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
    childrenProperties: graphql`
        fragment ChildrenPropertiesList_childrenProperties on ChildrenGroup@relay(plural: true) {
            type {
                name
                properties {
                    name
                }
            }
            objects {
                id
                name
                properties {
                    id
                    value
                }
            }
        }
    `
});

