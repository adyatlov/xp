import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';


ReactDOM.render(<App />, document.getElementById("root"));

// For debug purposes. Use it instead log.
console.shallowCloneLog = function(){
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

