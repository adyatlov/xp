let req = new XMLHttpRequest()

req.onload = (e) => {
    console.log(req.responseText)
};

window.onload = () => {
    req.open("GET", "http://localhost:7777/api/objects/cluster");
    req.send();
};
