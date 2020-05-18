import React from "react";
import {QueryRenderer} from 'react-relay';
import graphql from 'babel-plugin-relay/macro';

import environment from "./relayEnvironment";

export default class Plugin extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            url: ''
        };
        this.handleURLChange = this.handleURLChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }
    handleURLChange(event) {
        this.setState({url: event.target.value});
    }
    handleSubmit(event) {
        alert('A name was submitted: ' + this.state.url);
        event.preventDefault();
    }
    render() {
        const {url} = this.state
        return (
            <QueryRenderer
                environment={environment}
                query={graphql`query PluginQuery($url: String) {plugins(url: $url) {name}}`}
                variables={{url}}
                render={({error, props}) => {
                    if (error) {
                        return <div>Error!</div>;
                    }
                    let plugins = null
                    if (props) {
                        plugins = props.plugins
                    }
                    return (
                        <div className="btn-group" role="group">
                            <div className="input-group">
                                <input value={this.state.url} onChange={this.handleURLChange}
                                       type="text" className="form-control" placeholder="Insert dataset URL"
                                       aria-label="Dataset URL" aria-describedby="btnGroupAddon"/>
                                <PluginDropdown plugins={plugins}/>
                                <div className="input-group-append">
                                    <button type="button" className="btn btn-dark text-nowrap">Open</button>
                                </div>
                            </div>
                        </div>
                    );
                }}
            />
        );
    }
}

function PluginDropdown(props) {
    if (!props.plugins) {
        return (
            <div className="input-group-append dropdown">
                <button type="button" className="btn btn-secondary" disabled>
                    Loading plugins...
                </button>
                <div className="dropdown-menu" aria-labelledby="pluginsDropdown">Loading plugins...</div>
            </div>
        )
    }
    if (props.plugins.length === 0) {
        return (
            <div className="input-group-append dropdown">
                <button  type="button" className="btn btn-secondary" disabled>
                    Plugins not found
                </button>
                <div className="dropdown-menu" aria-labelledby="pluginsDropdown">Not found</div>
            </div>
            )
    }
    return (
        <div className="input-group-append dropdown">
            <button id="pluginsDropdown" type="button" className="btn btn-secondary dropdown-toggle"
                    data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                Choose plugin
            </button>
            <div className="dropdown-menu" aria-labelledby="pluginsDropdown">
                {props.plugins.map((plugin, index) => {
                    return <a className="dropdown-item" onClick={props.onSelected} key={index} href="#">{plugin.name}</a>
                })}
            </div>
        </div>
    );
}