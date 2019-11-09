import React from 'react';
import registry from './registry'

function MetricTypeName(props) {
    let n = registry.metricTypes[props.name].displayName;
    return (
        <span className="metric-type-name">{n}</span>
    );
}

export default MetricTypeName;
