import React from 'react';
import ReactDOM from 'react-dom';
import {BrowserRouter as Router, Route} from "react-router-dom";

import App from './App';

// For debug purposes. Use it instead log.
console.myLog = function(){
    let typeString = Function.prototype.call.bind(Object.prototype.toString)
    console.log.apply(console, Array.prototype.map.call(arguments, function(x){
        switch (typeString(x).slice(8, -1)) {
            case 'Number': case 'String': case 'Undefined': case 'Null': case 'Boolean': return x;
            case 'Array': return x.slice();
            default:
                let out = Object.create(Object.getPrototypeOf(x));
                out.constructor = x.constructor;
                for (let key in x) {
                    out[key] = x[key];
                }
                Object.defineProperty(out, 'constructor', {value: x.constructor});
                return out;
        }
    }));
}

ReactDOM.render(
    <Router>
        <Route path={[
            "/o/:datasetId/:objectId/:childrenTypeName",
            "/o/:datasetId/:objectId",
            "/"]} component={App}/>
    </Router>,
    document.getElementById('root')
);



