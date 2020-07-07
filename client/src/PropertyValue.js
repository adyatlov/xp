import React from 'react';

export default function PropertyValue(props) {
    let {value, type} = props;
    if (type === "TIMESTAMP") {
        value = formatDate(new Date(parseInt(value)));
    }
    return (
        <>{value}</>
    )
}

function formatDate(date) {
    function pad(number) {
        if (number < 10) {
            return '0' + number;
        }
        return number;
    }
    return date.getUTCFullYear() +
        '-' + pad(date.getUTCMonth() + 1) +
        '-' + pad(date.getUTCDate()) +
        ' ' + pad(date.getUTCHours()) +
        ':' + pad(date.getUTCMinutes()) +
        ':' + pad(date.getUTCSeconds());
}
