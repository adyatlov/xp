import React from 'react';

function ObjectTypeName(props) {
    return (
        <span>{props.name}</span>
    );
}

export function ObjectTypePlural(props) {
    return (
        <span className="object-type-name">{props.name}</span>
        // <span className="object-type-name">{registry.objectTypes[props.name].pluralDisplayName}</span>
    );
}

export default ObjectTypeName;
