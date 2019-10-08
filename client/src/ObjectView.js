import React from 'react'
import ObjectSummaryPanel from "./ObjectSummaryPanel";
import ObjectTable from "./ObjectTable";

class ObjectView extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            selectedChildrenGroupType: null
        };
        const o = props.object;
        if (typeof o.children !== "undefined" && o.children.length !== 0) {
            this.state.selectedChildrenGroupType = o.children[0].type
        }
    }

    render() {
        const o = this.props.object;
        const t = this.state.selectedChildrenGroupType;
        let children = null;
        if (t !== null) {
            children = childrenByType(o, t);
        }
        return (
            <div className="row">
                <div className="col-3">
                    <ObjectSummaryPanel object={o} selectedChildrenGroupType={t}/>
                </div>
                <div className="col-9">
                    {children !== null &&
                    <ObjectTable objects={children} handleSelectObject={this.props.handleSelectObject}/>
                    }
                </div>
            </div>
        );
    }
}

function childrenByType(object, type) {
    if (typeof object.children === "undefined") {
        return null;
    }
    for (let i = 0; i < object.children.length; i++) {
        if (object.children[i].type === type) {
            return object.children[i].objects
        }
    }
    return null;
}

export default ObjectView;
