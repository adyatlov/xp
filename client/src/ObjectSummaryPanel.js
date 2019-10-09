import React from "react";
import ObjectTypeName, {ObjectTypePlural} from "./ObjectType";
import MetricTypeName from "./MetricType";

function ObjectSummaryPanel(props) {
    const o = props.object;
    const t = props.selectedChildrenGroupType;
    const hasChildren = typeof o.children !== "undefined" && o.children.length !== 0;
    return (
        <div className="card">
            <div className="card-header list-group-item-action text-white bg-info">
                <><strong><ObjectTypeName name={o.type}/>:</strong> {o.name}</>
            </div>
            <MetricList metrics={o.metrics}/>
            {hasChildren &&
            <ObjectGroupList objectGroups={o.children} selectedChildrenGroupType={t}/>
            }
        </div>
    )
}

function MetricList(props) {
    const metrics = props.metrics
    return (
        <div className="card-body">
            {metrics.map(m => (
                <span key={m.type}>
                    <strong><MetricTypeName name={m.type}/>: </strong>{m.value}<br/>
                </span>
            ))}
        </div>
    );
}

function ObjectGroupList(props) {
    const gg = props.objectGroups;
    const t = props.selectedChildrenGroupType;
    return (
        <ul className="list-group list-group-flush">
            {gg.map(g => (
                <ChildrenGroupItem key={g.type} selected={g.type === t} group={g}/>
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
    return (
        <li className={tcl} href="#"><ObjectTypePlural name={props.group.type}/>
            <span className={ccl}>{props.group.objects.length}</span>
        </li>
    );
}

export default ObjectSummaryPanel