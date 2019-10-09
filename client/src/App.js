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
            isLoaded: false,
            cluster: null,
            object: null,
            error: null
        }
    }

    setStateAndCheck(s) {
        if (s !== null) {
            this.setState(s);
        }
        if (this.state.cluster !== null && registry.objectTypes !== null && registry.metricTypes !== null) {
            this.setState({isLoaded: true})
        }
    }

    handleSelectObject(e) {
        let t = e.target.dataset.otype;
        let id = e.target.dataset.oid;
        fetch(endpointObjects + "/" + t + "/" + id)
            .then(res => res.json())
            .then(
                (result) => {
                    this.setState({object: result});
                },
                (error) => {
                    console.error(error);
                }
            );
        e.preventDefault();
    }

    componentDidMount() {
        fetch(endpointCluster)
            .then(res => res.json())
            .then(
                (result) => {
                    this.setStateAndCheck({cluster: result, object: result});
                },
                (error) => {
                    this.setState({
                        isLoaded: true,
                        error: error
                    });
                }
            );
        fetch(endpointObjectTypes)
            .then(res => res.json())
            .then(
                (result) => {
                    registry.objectTypes = result;
                    this.setStateAndCheck(null);
                },
                (error) => {
                    this.setState({
                        error: error
                    });
                }
            );
        fetch(endpointMetricTypes)
            .then(res => res.json())
            .then(
                (result) => {
                    registry.metricTypes = result;
                    this.setStateAndCheck(null);
                },
                (error) => {
                    this.setState({
                        error: error
                    });
                }
            );
    }

    render() {
        const isLoaded = this.state.isLoaded;
        const cluster = this.state.cluster;
        const object = this.state.object;
        const error = this.state.error;
        if (error) {
            return <div>Error: {error.message}</div>;
        } else if (!isLoaded) {
            return <div>Loading...</div>;
        }
        return (
            <>
                <ClusterBar clusterName={cluster.name}/>
                <ObjectView object={object} handleSelectObject={this.handleSelectObject}/>
            </>
        )
    }
}

export default App;
