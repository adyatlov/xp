import React from "react";
import {QueryRenderer, commitMutation} from 'react-relay';
import graphql from 'babel-plugin-relay/macro';
import environment from "./relayEnvironment";


const query = graphql`
    query DatasetAdderQuery ($url: String) {
        compatiblePlugins(url: $url) {
            name
        }
    }
`;

const mutation = graphql`
    mutation DatasetAdderAddDatasetMutation($pluginName: String!, $url: String!) {
        addDataset(pluginName: $pluginName, url: $url) {
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
                pluginName: pluginName,
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
                        plugins = props.compatiblePlugins;
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
        const plugins = this.props.plugins;
        const pluginName = this.state.pluginName;
        // Do nothing when the plugins array is not set on loading.
        if (!plugins) {
            return;
        }
        // If only one compatible plugin was found then choose it straightaway.
        if (plugins.length === 1 && pluginName !== plugins[0].name) {
            this.setState({pluginName: plugins[0].name})
            return;
        }
        // If the current chosen plugin is not in the list of compatible plugins then reset it.
        if (pluginName === null) {
            return;
        }
        const compatible = plugins.some((plugin) => {
            return plugin.name === pluginName
        })
        if (!compatible) {
            this.setState({pluginName: null})
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
        const openDisabled = (url === "" || !plugins || plugins.length === 0 || !pluginName);
        return (
            <div className="btn-toolbar d-flex" role="toolbar" aria-label="Toolbar for opening datasets">
                <div className="input-group flex-grow-1">
                    <UrlInput url={url} onChange={this.handleURLChange}/>
                    {url !=="" &&
                    <PluginSelector pluginName={pluginName}
                                                  plugins={plugins}
                                                  onChange={this.handlePluginNameChange}/>}
                </div>
                <button onClick={this.handleAddDataset}
                        type="button" disabled={openDisabled}
                        className="btn btn-primary text-nowrap ml-2">Open</button>
            </div>
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
    let handleOnClick = (event) => {
        event.preventDefault();
        onChange(event.target.value);
    }
    const {onChange, pluginName, plugins} = props;
    let disabled = true;
    let text = pluginName;
    (() => {
        if (!plugins) {
            text = (
                <>
                    <span className="spinner-grow spinner-grow-sm mr-2" role="status" aria-hidden="true"/>
                    Loading plugins...
                </>
            );
            return;
        }
        if (plugins.length === 0) {
            text = "No compatible plugins found";
            return;
        }
        if (plugins.length > 1) {
            disabled = false;
        }
        if (!pluginName) {
            text = "Choose plugin";
        }
    })();
    let className = "btn btn-primary dropdown-toggle";
    if (disabled) {
        className += " dropdown-toggle-arrow-off";
    }
    return (
        <div className="input-group-append">
            <button id="pluginsDropdown" disabled={disabled}
                    className={className}
                    type="button"
                    data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                {text}
            </button>
            <div className="dropdown-menu" aria-labelledby="pluginsDropdown">
                {plugins && plugins.map((plugin, index) => {
                    return <button onClick={handleOnClick}
                                   key={index}
                                   value={plugin.name}
                                   className="dropdown-item">{plugin.name}</button>
                })}
            </div>
        </div>
    );
}

