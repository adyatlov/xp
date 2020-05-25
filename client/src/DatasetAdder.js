import React from "react";
import {QueryRenderer} from 'react-relay';
import graphql from 'babel-plugin-relay/macro';
import environment from "./relayEnvironment";


const query = graphql`
    query DatasetAdderQuery ($url: String) {
        plugins(url: $url) {
            name
        }
    }
`;

export default class DatasetAdderQuery extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            url: "",
            pluginName: null
        };
        this.handleURLChange = this.handleURLChange.bind(this);
        this.handlePluginNameChange = this.handlePluginNameChange.bind(this);
    }

    handleURLChange(url) {
        this.setState({url: url})
    }

    handlePluginNameChange(pluginName) {
        this.setState( {pluginName: pluginName})
    }

    render() {
        const {url} = this.state;
        let {pluginName} = this.state;
        return(
            <QueryRenderer
                environment={environment}
                query={query}
                variables={{url}}
                // fetchPolicy={"store-and-network"}
                render={({error, props}) => {
                    if (error) {
                        console.log(error.text);
                    }
                    let plugins = null
                    if (props) {
                        plugins = props.plugins;
                        if (plugins.length === 1) {
                            pluginName = plugins[0].name;
                        }
                    }
                    return (
                        <DatasetAdder url={url}
                                      pluginName={pluginName}
                                      plugins={plugins}
                                      onURLChange={this.handleURLChange}
                                      onPluginNameChange={this.handlePluginNameChange}/>
                    );
                }}/>
        );
    }
}


function DatasetAdder (props) {
        const {url, plugins, onURLChange, onPluginNameChange} = props
        let {pluginName} = props
        if (!plugins) {
            return (
                <InputGroup>
                    <UrlInput url={url} onChange={onURLChange}/>
                    <Message>
                        <span className="spinner-grow spinner-grow-sm mr-2" role="status" aria-hidden="true"/>
                        Loading plugins...
                    </Message>
                    <OpenButton disabled/>
                </InputGroup>
            );
        }
        if (plugins.length === 0) {
            return (
                <InputGroup>
                    <UrlInput url={url} onChange={onURLChange}/>
                    <Message>No compatible plugins found</Message>
                    <OpenButton disabled/>
                </InputGroup>
            );
        }
        const openDisabled = (url === "" || plugins.length === 0);
        return (
            <InputGroup>
                <UrlInput url={url} onChange={onURLChange}/>
                <PluginSelector pluginName={pluginName} plugins={plugins} onChange={onPluginNameChange}/>
                <OpenButton disabled={openDisabled}/>
            </InputGroup>
        );
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
            <button disabled={props.disabled} type="button" className="btn btn-secondary text-nowrap">Open</button>
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

// const mutation = graphql`
//     mutation DatasetOpenerAddDatasetMutation($plugin: String!, $url: String!) {
//         addDataset(plugin: $plugin, url: $url) {
//             id
//             root {
//                 type {
//                     name
//                 }
//                 name
//             }
//         }
//     }
// `

// function addDataset(environment, url, pluginName) {
//     return commitMutation(
//         environment,
//         {
//             mutation,
//             variables: {
//                 input: {url, pluginName}
//             },
//             onError: (error) => {
//                 console.error(error.message);
//             }
//         }
//     )
// }

