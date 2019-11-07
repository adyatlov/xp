import React from 'react';
import ClusterBar from "./ClusterBar";
import ObjectView from "./ObjectView";
import registry from "./registry";

const serverURL = "http://localhost:7777/api";
const endpointCluster = serverURL + "/objects/cluster";
const endpointObjects = serverURL + "/objects";
// const endpointMetrics = serverURL + "/metrics";
const endpointObjectTypes = serverURL + "/objectTypes";
const endpointMetricTypes = serverURL + "/metricTypes";

class App extends React.Component {
    constructor(props) {
        super(props);
        this.handleSelectObject = this.handleSelectObject.bind(this);
        this.state = {
            cluster: null,
            object: null,
            error: null
        }
    }

    componentDidMount() {
        console.log("componentDidMount");
        this.update();
    }

    componentDidUpdate() {
        console.log("componentDidUpdate");
        this.update();
    }

    // TODO: delete
    handleSelectObject(e) {
        this.update();
        e.preventDefault();
    }

    update() {
        this.updateRegistry();
        this.updateCluster();
        this.updateObject();
    }

    updateCluster() {
        if (this.state.cluster !== null) {
            return;
        }
        fetch(endpointCluster)
            .then(res => res.json())
            .then(
                (result) => {
                    this.setState({cluster: result});
                    console.log("Cluster updated")
                },
                (error) => {
                    console.error(error);
                }
            );
    }

    updateObject() {
        const match = this.props.match;
        let endpoint = endpointCluster;
        if (match.path === "/") {
            const object = this.state.object;
            if (object && object.type === "cluster") {
                return;
            }
            endpoint = endpointCluster;
        } else if (match.path === "/o/:type/:id") {
            const object = this.state.object;
            const t = match.params.type;
            const id = match.params.id;
            if (object && object.type === t && object.id === id) {
                return;
            }
            endpoint = endpointObjects + "/" + t + "/" + id;
        } else {
            return;
        }
        fetch(endpoint)
            .then(res => res.json())
            .then(
                (result) => {
                    this.setState({object: result});
                    console.log("Object updated");
                },
                (error) => {
                    console.error(error);
                }
            );
    }

    updateRegistry() {
        if (!registry.objectTypes) {
            fetch(endpointObjectTypes)
                .then(res => res.json())
                .then(
                    (result) => {
                        registry.objectTypes = result;
                        this.setState({})
                    },
                    (error) => {
                        this.setState({
                            error: error
                        });
                    }
                );
        }
        if (!registry.metricTypes) {
            fetch(endpointMetricTypes)
                .then(res => res.json())
                .then(
                    (result) => {
                        registry.metricTypes = result;
                        this.setState({});
                    },
                    (error) => {
                        this.setState({
                            error: error
                        });
                    }
                );
        }
    }

    render() {
        console.log("RENDER", this.state);
        const cluster = this.state.cluster;
        const object = this.state.object;
        const error = this.state.error;
        if (error) {
            return <div>Error: {error.message}</div>;
        } else if (registry.objectTypes && registry.metricTypes && this.state.cluster && this.state.object) {
            return (
                <>
                    <ClusterBar clusterName={cluster.name}/>
                    <ObjectView object={object} handleSelectObject={this.handleSelectObject}/>
                </>);
        }
        return <div>Loading...</div>;
    }

}

export default App;
