import React from 'react';
import ObjectTypeName from "./ObjectType";
import MetricTypeName from "./MetricType";
import ObjectLink from "./ObjectLink";

function ObjectTable(props) {
    const objects = props.objects;
    const o = objects[0];
    return (
        <table className="table table-bordered">
            <thead>
            <tr>
                <th scope="col">
                    <ObjectTypeName name={o.type}/>
                </th>
                {o.metrics.map(m => (
                    <th key={m.type} scope="col">
                        <MetricTypeName name={m.type}/>
                    </th>
                ))}
            </tr>
            </thead>
            <tbody>
            {objects.map(obj => (
                <tr key={obj.id}>
                    <td>
                        <ObjectLink handleSelectObject={props.handleSelectObject} object={obj}/>
                    </td>
                    {obj.metrics.map(m => (
                        <td key={m.type}>{m.value}</td>
                    ))}
                </tr>
            ))}
            </tbody>
        </table>
    );
}

export default ObjectTable;

