import React from 'react';
import ObjectTypeName from "./ObjectType";
import MetricTypeName from "./MetricType";
import ObjectLink from "./ObjectLink";

class ObjectTable extends React.Component {
    render(props) {
        const o = props.objects;
        if (!o) {
            return (
                <div>No objects</div>
            );
        }
        return (
            <table className="table table-bordered">
                <thead>
                <tr>
                    <th scope="col">
                        <ObjectTypeName name={props.typeName}/>
                    </th>
                    {o[0].metrics.map(m => (
                        <th key={m.typeName} scope="col">
                            <MetricTypeName name={m.typeName}/>
                        </th>
                    ))}
                </tr>
                </thead>
                <tbody>
                {o.map(obj => (
                    <tr key={obj.objectId}>
                        <td>
                            <ObjectLink typeName={props.typeName} objectId={obj.objectId}>{obj.name}</ObjectLink>
                        </td>
                        {obj.metrics.map(m => (
                            <td key={m.typeName}>{m.value}</td>
                        ))}
                    </tr>
                ))}
                </tbody>
            </table>
        );
    }
}

export default ObjectTable;

function childrenByType(object, typeName) {
    if (!object || !object.children || !typeName) {
        return null;
    }
    for (let i = 0; i < object.children.length; i++) {
        if (object.children[i].typeName === typeName) {
            return object.children[i].objects
        }
    }
    return null;
}
