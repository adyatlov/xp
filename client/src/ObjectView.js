import React from 'react'
import ObjectSummaryPanel from "./ObjectSummaryPanel";
import ObjectTable from "./ObjectTable";

class ObjectView extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            object: props.object,
            selectedChildrenGroupType: "",
        };
        if (typeof this.state.object.children !== "undefined" && this.state.object.children.length !== 0) {
            this.state.selectedChildrenGroupType =  this.state.object.children[0].type
        }
    }

    render() {
        const object = this.state.object;
        const selectedChildrenGroupType = this.state.selectedChildrenGroupType;
        let children = null;
        if (selectedChildrenGroupType !== "") {
            children = childrenByType(object, selectedChildrenGroupType);
        }
        return (
            <div className="row">
                <div className="col-2">
                    <ObjectSummaryPanel object={object} selectedChildrenGroupType={selectedChildrenGroupType}/>
                </div>
                <div className="col-10">
                    {children !== null &&
                    <ObjectTable objects={children}/>
                    }
                </div>
            </div>
        );
    }
}

function childrenByType(object, type) {
    for (let i = 0; i < object.children.length; i++) {
        if (object.children[i].type === type) {
            return object.children[i].objects
        }
    }
    return [];
}

export default ObjectView;
