import React from "react";
import {QueryRenderer} from 'react-relay';
import graphql from 'babel-plugin-relay/macro';

import environment from "./relayEnvironment";

export default class Plugin extends React.Component {
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
        return (
            <div className="btn-group" role="group">
                <div className="input-group">
                    <input value={this.state.url} onChange={this.handleURLChange}
                           type="text" className="form-control" placeholder="Insert dataset URL"
                           aria-label="Dataset URL" aria-describedby="btnGroupAddon"/>
                    <PluginSelector url={this.state.url} onSelect={this.handlePluginNameChange}/>
                    <div className="input-group-append">
                        <button type="button" className="btn btn-dark text-nowrap">Open</button>
                    </div>
                </div>
            </div>
        );
    }
}

class PluginSelector extends React.Component {
    constructor(props) {
        super(props);
        this.handleChange = this.handleChange.bind(this);
        this.state = {
            pluginName: "Choose plugin"
        }
        this.singlePluginElement = null;
        this.setSinglePluginElementRef = element => {
            this.singlePluginElement = element;
        }
    }

    handleChange(event) {
        event.preventDefault();
        this.setState({pluginName: event.target.value})
        this.props.onSelect(event.target.value)
    }

    componentDidMount() {
        if (this.singlePluginElement) {
            this.singlePluginElement.click();
        }
    }

    render() {
        const {url} = this.props
        let message = function(text) {
            return (
                <div className="input-group-append dropdown">
                    <button type="button" className="btn btn-secondary" disabled>
                        {text}
                    </button>
                    <div className="dropdown-menu" aria-labelledby="pluginsDropdown">{text}</div>
                </div>
            );
        }
        return(
            <QueryRenderer
                environment={environment}
                query={graphql`query PluginQuery($url: String) {plugins(url: $url) {name}}`}
                variables={{url}}
                render={({error, props}) => {
                    if (error) {
                        return <div>Error: {error}</div>;
                    }
                    let plugins
                    if (props) {
                        plugins = props.plugins;
                    }
                    if (!plugins) {
                        return message("Loading plugins...");
                    }
                    if (plugins.length === 0) {
                        return message("No compatible plugins found");
                    }
                    const {pluginName} = this.state
                    let refAttr = {}
                    if  (plugins.length === 1) {
                        refAttr.ref = this.setSinglePluginElementRef
                    }
                    return (
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
                }}
            />
        );
    }
}