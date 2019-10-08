import React from 'react';

function ObjectTable(props) {
    return (
        <table className="table table-bordered">
            <tbody>
            {props.objects.map(obj => (
                <tr key={obj.id}>
                    <td>
                        {obj.name}
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

