import React from 'react';
import ClusterBar from "./ClusterBar";
import ObjectView from "./ObjectView";

class App extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            isLoaded: false,
            cluster: null,
            error: null
        }
    }

    componentDidMount() {
        fetch("http://localhost:7777/api/objects/cluster")
            .then(res => res.json())
            .then(
                (result) => {
                    this.setState({
                        isLoaded: true,
                        cluster: result
                    });
                },
                (error) => {
                    this.setState({
                        isLoaded: true,
                        error: error
                    });
                }
            )
    }

    render() {
        const {isLoaded, cluster, error} = this.state;
        if (error) {
            return <div>Error: {error.message}</div>;
        } else if (!isLoaded) {
            return <div>Loading...</div>;
        }
        return (
            <>
                <ClusterBar clusterName={cluster.name}/>
                <ObjectView object={cluster}/>
            </>
        )
    }
}

export default App;
