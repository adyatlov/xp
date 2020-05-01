import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import {BrowserRouter as Router, Route} from "react-router-dom";

ReactDOM.render(
    <Router>
        <Route path={["/o/:typeName/:objectId/:childrenTypeName", "/o/:typeName/:objectId", "/"]} component={App}/>
    </Router>,
    document.getElementById('root'));
