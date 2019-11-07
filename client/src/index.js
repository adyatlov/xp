import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import {BrowserRouter as Router, Route} from "react-router-dom";

ReactDOM.render(
    <Router>
        <Route path={["/o/:type/:id", "/"]} component={App}/>
    </Router>,
    document.getElementById('root'));
