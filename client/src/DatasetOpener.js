import React from "react";
import {QueryRenderer} from 'react-relay';
import graphql from 'babel-plugin-relay/macro';

import environment from "./relayEnvironment";

const query = graphql`
    query PluginQuery($url: String) {
        plugins(url: $url) {
            name
        }
    }`

export default class DatasetOpener extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            url: "",
            pluginName: null,
        };
        this.handleURLChange = this.handleURLChange.bind(this);
        this.handlePluginNameChange = this.handlePluginNameChange.bind(this);
    }
    handleURLChange(event) {
        this.setState({url: event.target.value});
    }
    handlePluginNameChange(pluginName) {
        this.setState({pluginName: pluginName});
    }
    render() {
        const {url} = this.state
        let {pluginName} = this.state
        return (
            <QueryRenderer
                environment={environment}
                query={query}
                variables={{url}}
                render={({error, props}) => {
                    if (error) {
                        return (
                            <InputGroup>
                                <UrlInput url={url} onChange={this.handleURLChange}/>
                                <Message>{error.message}</Message>
                                <OpenButton disabled={true}/>
                            </InputGroup>
                        );
                    }
                    let plugins;
                    if (props) {
                        plugins = props.plugins;
                    }
                    if (!plugins) {
                        return (
                            <InputGroup>
                                <UrlInput url={url} onChange={this.handleURLChange}/>
                                <Message>Loading plugins...</Message>
                                <OpenButton disabled={true}/>
                            </InputGroup>
                        );
                    }
                    if (plugins.length === 0) {
                        return (
                            <InputGroup>
                                <UrlInput url={url} onChange={this.handleURLChange}/>
                                <Message>No compatible plugins found</Message>
                                <OpenButton disabled={true}/>
                            </InputGroup>
                        );
                    }
                    if (pluginName == null) {
                        pluginName = "Choose plugin";
                    }
                    let openDisabled = (url === "" || plugins.length === 0);
                    return (
                        <InputGroup>
                            <UrlInput url={url} onChange={this.handleURLChange}/>
                            <PluginSelector pluginName={pluginName} plugins={plugins} onSelect={this.handlePluginNameChange}/>
                            <OpenButton disabled={openDisabled}/>
                        </InputGroup>
                    );
                }}
            />
        );
    }
}

function InputGroup(props) {
    return(
        <div className="btn-group" role="group">
            <div className="input-group">
                {props.children}
            </div>
        </div>
    );
}

function UrlInput(props) {
    return(
        <input value={props.url} onChange={props.onChange}
               type="text" className="form-control" placeholder="Insert dataset URL"
               aria-label="Dataset URL" aria-describedby="btnGroupAddon"/>
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

function OpenButton(props) {
    return(
        <div className="input-group-append">
            <button disabled={props.disabled} type="button" className="btn btn-dark text-nowrap">Open</button>
        </div>
    );
}

class PluginSelector extends React.Component {
    constructor(props) {
        super(props);
        this.handleChange = this.handleChange.bind(this);
        this.singlePluginElement = null;
        this.setSinglePluginElementRef = element => {
            this.singlePluginElement = element;
        }
    }

    componentDidMount() {
        if (this.singlePluginElement != null)  {
            this.singlePluginElement.click();
        }
    }

    handleChange(event) {
        event.preventDefault();
        this.setState({pluginName: event.target.value})
        this.props.onSelect(event.target.value)
    }

    render() {
        const {pluginName, plugins} = this.props
        let refAttr = {}
        if  (plugins.length === 1) {
            refAttr.ref = this.setSinglePluginElementRef
        }
        return(
            <div className="input-group-append dropdown">
                <button id="pluginsDropdown" type="button" className="btn btn-secondary dropdown-toggle"
                        data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                    {pluginName}
                </button>
                <div className="dropdown-menu" aria-labelledby="pluginsDropdown">
                    {plugins.map((plugin, index) => {
                        return <button onClick={this.handleChange}
                                       key={index}
                                       value={plugin.name}
                                       {...refAttr}
                                       className="dropdown-item">{plugin.name}</button>
                    })}
                </div>
            </div>
        );
    }
}


