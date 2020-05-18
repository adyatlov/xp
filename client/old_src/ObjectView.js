import React from 'react';
import graphql from 'babel-plugin-relay/macro';
import {createFragmentContainer} from "react-relay";
import {Route, Switch, withRouter} from "react-router-dom";
import ObjectSummaryPanel from "./ObjectSummaryPanel";
import ObjectTable from "./ObjectTable";


class ObjectView extends React.Component {
    componentDidMount() {
        const object = this.props.object;
        if (!object) {
            return
        }
        if (this.props.match.path !== "/") {
            return;
        }
        const typeName = object.typeName;
        const objectId = object.objectId;
        const history = this.props.history;
        history.replace("/o/" + typeName + "/" + objectId);
    }

    render() {
        const object = this.props.object;
        if (!object) {
            return (
                <div>Not found</div>
            );
        }
        return (
            <div className="row">
                <div className="col-3">
                    <ObjectSummaryPanel object={object}/>
                </div>
                <div className="col-9">
                    <Switch>
                        <Route exact path="/o/:typeName/:objectId/:childrenTypeName">
                            <ObjectTable/>
                        </Route>
                        <Route exact path="/o/:typeName/:objectId">
                            <div>Metrics go here</div>
                        </Route>
                    </Switch>
                </div>
            </div>
        )
    }
}

export default createFragmentContainer(
    withRouter(ObjectView),
    {
        object: graphql`
            fragment ObjectView_object on Object {
                typeName
                objectId
                name
                metrics {
                    typeName
                    value
                }
                children {
                    typeName
                    objects {
                        name
                        objectId
                        metrics {
                            typeName
                            value
                        }
                    }
                }
            }
        `
    }
);


