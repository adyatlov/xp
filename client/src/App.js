import React from 'react';

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
                // Note: it's important to handle errors here
                // instead of a catch() block so that we don't swallow
                // exceptions from actual bugs in components.
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
        } else {
            return (
                <ul>
                    {cluster.children.map(children => (
                        <li key={children.type}>
                            {children.type}: {children.objects.length}
                            <ul>
                                {children.objects.map(obj => (
                                    <li key={obj.id}>
                                        {obj.name}: {obj.metrics.map(m => (
                                        <span>{m.type}: {m.value}</span>
                                    ))}
                                    </li>
                                    ))}
                            </ul>
                        </li>
                    ))}
                </ul>
            );
        }
    }
}

export default App;
