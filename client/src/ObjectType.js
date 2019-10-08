import React from 'react';
import registry from './registry'

function ObjectTypeName(props) {
    return (
        <span className={props.className}>{registry.objectTypes[props.name].displayName}</span>
    );
}

export function ObjectTypePlural(props) {
    return (
        <span className="object-type-name">{registry.objectTypes[props.name].pluralDisplayName}</span>
    );
}

export default ObjectTypeName;
