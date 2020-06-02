import React from "react";
import {QueryRenderer, commitMutation} from 'react-relay';
import graphql from 'babel-plugin-relay/macro';
import environment from "./relayEnvironment";


const query = graphql`
    query DatasetAdderQuery ($url: String) {
        plugins(url: $url) {
            name
        }
    }
`;

const mutation = graphql`
    mutation DatasetAdderAddDatasetMutation($plugin: String!, $url: String!) {
        addDataset(plugin: $plugin, url: $url) {
            id
        }
    }
`
function addDataset(environment, pluginName, url) {
    return commitMutation(
        environment,
        {
            mutation,
            variables: {
                plugin: pluginName,
                url: url
            },
            onError: (error) => {
                console.error(error.message);
            },
        }
    )
}

export default class DatasetAdderQuery extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            url: "",
        };
        this.handleURLChange = this.handleURLChange.bind(this);
    }

    handleURLChange(url) {
        this.setState({url: url})
    }

    render() {
        const {url} = this.state;
        return(
            <QueryRenderer
                environment={environment}
                query={query}
                variables={{url}}
                fetchPolicy={"store-and-network"}
                render={({error, props}) => {
                    if (error) {
                        console.error(error.text);
                    }
                    let plugins = null
                    if (props) {
                        plugins = props.plugins;
                    }
                    return (
                        <DatasetAdder plugins={plugins} onURLChange={this.handleURLChange}/>
                    );
                }}/>
        );
    }
}

class DatasetAdder extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            url: "",
            pluginName: null,
        }
        this.handleURLChange = this.handleURLChange.bind(this);
        this.handlePluginNameChange = this.handlePluginNameChange.bind(this);
        this.handleAddDataset = this.handleAddDataset.bind(this);
    }

    componentDidUpdate(prevProps, prevState) {
        if (this.props.plugins &&
            this.props.plugins.length === 1 &&
            this.state.pluginName !== this.props.plugins[0].name) {
            this.setState({pluginName: this.props.plugins[0].name})
        }
    }

    handleURLChange(url) {
        this.setState({url: url});
        this.props.onURLChange(url);
    }

    handlePluginNameChange(pluginName) {
        this.setState({pluginName: pluginName});
    }

    handleAddDataset() {
        addDataset(environment, this.state.pluginName, this.state.url);
    }

    render() {
        const {plugins} = this.props;
        const {url, pluginName} = this.state;
        const openDisabled = (url === "" || !plugins || plugins.length === 0);
        let selector = (
            <PluginSelector pluginName={pluginName} plugins={plugins} onChange={this.handlePluginNameChange}/>
        );
        if (!plugins) {
            selector = (
                <Message>
                    <span className="spinner-grow spinner-grow-sm mr-2" role="status" aria-hidden="true"/>
                    Loading plugins...
                </Message>
            );
        } else if (plugins.length === 0) {
            selector = (
                <Message>No compatible plugins found</Message>
            );
        }
        return (
            <InputGroup>
                <UrlInput url={url} onChange={this.handleURLChange}/>
                {selector}
                <OpenButton onAddDataset={this.handleAddDataset} disabled={openDisabled}/>
            </InputGroup>
        );
    }
}

function UrlInput(props) {
    const {url, onChange} = props;
    let handleOnChange = (event) => {
        onChange(event.target.value);
    }
    return(
        <input value={url} onChange={handleOnChange}
               type="text" className="form-control" placeholder="Insert dataset URL"
               aria-label="Dataset URL" aria-describedby="btnGroupAddon"/>
    );
}

function PluginSelector(props) {
    const {plugins, onChange} = props;
    let {pluginName} = props
    if (!pluginName) {
        pluginName = "Choose plugin";
    }
    let handleOnClick = (event) => {
        event.preventDefault();
        onChange(event.target.value);
    }
    return(
        <div className="input-group-append dropdown">
            <button id="pluginsDropdown" type="button" className="btn btn-secondary dropdown-toggle"
                    data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                {pluginName}
            </button>
            <div className="dropdown-menu" aria-labelledby="pluginsDropdown">
                {plugins.map((plugin, index) => {
                    return <button onClick={handleOnClick}
                                   key={index}
                                   value={plugin.name}
                                   className="dropdown-item">{plugin.name}</button>
                })}
            </div>
        </div>
    );
}

function OpenButton(props) {
    return(
        <div className="input-group-append">
            <button onClick={props.onAddDataset}
                    type="button" disabled={props.disabled}
                    className="btn btn-secondary text-nowrap">Open</button>
        </div>
    );
}

function InputGroup(props) {
    return(
        <div className="input-group">
            {props.children}
        </div>
    );
}

function Message(props) {
    return (
        <div className="input-group-append dropdown">
            <button type="button" className="btn btn-secondary" disabled>
                {props.children}
            </button>
            <div className="dropdown-menu" aria-labelledby="pluginsDropdown">placeholder</div>
        </div>
    );
}

