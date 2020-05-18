import React from "react";
import ObjectTypeName, {ObjectTypePlural} from "./ObjectType";
import MetricTypeName from "./MetricType";
import ChildrenLink from "./ChildrenLink";

class ObjectSummaryPanel extends React.Component {
    render() {
        const o = this.props.object;
        const t = this.props.childrenTypeName;
        const hasChildren = typeof o.children !== "undefined" && o.children.length !== 0;
        return (
            <div className="card">
                <div className="card-header list-group-item-action text-white bg-info">
                    <><strong><ObjectTypeName name={o.typeName}/>:</strong> {o.name}</>
                </div>
                <MetricList metrics={o.metrics}/>
                {hasChildren &&
                <ObjectGroupList objectGroups={o.children} childrenTypeName={t}/>
                }
            </div>
        )
    }
}

function MetricList(props) {
    const metrics = props.metrics;
    if (!metrics) {
        return (
            <div className="card-body">
                No metrics found
            </div>
        );
    }
    return (
        <div className="card-body">
            {metrics.map(m => (
                <span key={m.typeName}>
                    <strong><MetricTypeName name={m.typeName}/>: </strong>{m.value}<br/>
                </span>
            ))}
        </div>
    );
}

function ObjectGroupList(props) {
    const groups = props.objectGroups;
    return (
        <ul className="list-group list-group-flush">
            {groups.map(g => (
                <ChildrenGroupItem key={g.typeName} group={g}/>
            ))}
        </ul>
    );
}

function ChildrenGroupItem(props) {
    let tcl = "list-group-item list-group-item-action d-flex justify-content-between align-items-center";
    let ccl = "badge";
    if (props.selected) {
        tcl += " text-white bg-secondary";
        ccl += " badge-light";
    }
    const group = props.group;
    return (
        <li className={tcl}>
            <ChildrenLink childrenTypeName={group.typeName}>
                <ObjectTypePlural name={group.typeName}/>
            </ChildrenLink>
            <span className={ccl}>{group.objects.length}</span>
        </li>
    );
}


export default ObjectSummaryPanel;
