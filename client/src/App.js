import React from 'react';
import {useLocation} from "react-router-dom";
import {QueryRenderer} from 'react-relay';
import graphql from 'babel-plugin-relay/macro';

import environment from "./relayEnvironment";
import DatasetAdder from "./DatasetAdder";
// import Properties from "./Properties";
// import ChildrenTable from "./ChildrenTable";
// import {ObjectViewLayout} from "./ObjectViewLayout";

const query = graphql`
    query AppQuery {
        plugins {
            name
            description
        }
    }
`;

export default class App extends React.Component {
    render() {
        const {match} = this.props;
        return(
            <QueryRenderer
                environment={environment}
                query={query}
                // fetchPolicy={"store-and-network"}
                render={({error, props}) => {
                    if (error) {
                        return (<div>{error.text}</div>);
                    }
                    let plugins = null;
                    if (props) {
                        plugins = props.plugins
                    }
                    if (match.path === "/") {
                        return (
                            <>
                                <TopBar>
                                    <DatasetAdder/>
                                </TopBar>
                                <Test plugins={plugins}/>
                                <SelectDatasets/>
                            </>
                        );
                    }
                    return (
                        <>
                            <TopBar/>
                            <NoMatch/>
                        </>
                    );
                }
                } />
        );
    }
}

function Test(props) {
    if (!props.plugins) {
        return <div>Loading...</div>
    }
    return(
    <ul>
        {props.plugins.map((value, index) => {
            return(
                <li key={index}><h5>{value.name}</h5><p>{value.description}</p></li>
            );
        })}
    </ul>
    );
}

function TopBar(props) {
    return (
        <nav className="navbar navbar-light bg-light">
            <form className="form-inline">
                {props.children}
            </form>
            <span className="navbar-brand">XP</span>
        </nav>
    );
}

function MessageBox(props) {
    let className = "alert alert-light";
    if (props.warning) {
        className = "alert alert-warning";
    }
    return(
        <div id="root" className="container">
            <div className="row mt-5">
                <div className="col">
                    <div className={className} role="alert">
                        {props.children}
                    </div>
                </div>
            </div>
        </div>
    );
}

function SelectDatasets() {
    return (
        <MessageBox>
            <h4 className="alert-heading">Welcome to XP!</h4>
            <p>XP helps you to explore heterogeneous datasets uniformly.</p>
            <p>Please, select a dataset or open a new one:
                insert a dataset URL, choose one of the compatible plugins, and press "Open".</p>
        </MessageBox>
    );
}

function NoMatch() {
    let location = useLocation();
    return (
        <MessageBox warning>
            <h4 className="alert-heading">Error</h4>
            <p>Page <code>{location.pathname}</code> not found.</p>
        </MessageBox>
    );
}

// function LoadingSpinner() {
//    return(
//             <div style={{position: "fixed", top: "50%", left: "50%",
//             transform:"translate(-50%, -50%)"}}>
//                 <h4 className="text-secondary">
//                     Loading XP...
//                 </h4>
//                 <div className="text-center">
//                     <div className="spinner-grow text-secondary mt-3" style={{width: "2rem", height: "2rem"}} role="status"/>
//                 </div>
//             </div>
//    );
// }
