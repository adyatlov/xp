import React from 'react';

class ObjectSummaryPanel extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            object: props.object,
            selectedChildrenGroupType: props.selectedChildrenGroupType
        };
    };

    render() {
        const object = this.state.object;
        const selectedChildrenGroupType = this.state.selectedChildrenGroupType;
        return (
            <div className="card">
                <div className="card-header list-group-item-action text-white bg-info">{object.name}</div>
                <MetricList metrics={object.metrics}/>
                <ObjectGroupList objectGroups={object.children} selectedChildrenGroupType={selectedChildrenGroupType}/>
            </div>
        )
    }
}

function MetricList(props) {
    const metrics = props.metrics
    return (
        <div className="card-body">
            {metrics.map(m => (
                <span key={m.type}>
                    <strong>{m.type}: </strong>{m.value}<br/>
                </span>
            ))}
        </div>
    );
}

class ObjectGroupList extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            objectGroups: props.objectGroups,
            selectedChildrenGroupType: props.selectedChildrenGroupType
        };
    }

    render() {
        const objectGroups = this.state.objectGroups;
        const selectedChildrenGroupType = this.state.selectedChildrenGroupType;
        return (
            <ul className="list-group list-group-flush">
                {objectGroups.map(group => (
                    <ChildrenGroupItem key={group.type}
                                       selected={group.type === selectedChildrenGroupType}
                                       group={group}/>
                ))}
            </ul>
        );
    }
}

function ChildrenGroupItem(props) {
    let tcl = "list-group-item list-group-item-action d-flex justify-content-between align-items-center";
    let ccl = "badge";
    if (props.selected) {
        tcl += " text-white bg-secondary";
        ccl += " badge-light";
    }
    return (
        <li className={tcl} href="#">{props.group.type}
            <span className={ccl}>{props.group.objects.length}</span>
        </li>
    );
}

export default ObjectSummaryPanel