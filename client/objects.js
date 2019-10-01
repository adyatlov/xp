let req = new XMLHttpRequest()

req.onload = (e) => {
    let cluster = JSON.parse(req.responseText)
    console.log(cluster)
};

window.onload = () => {
    req.open("GET", "http://localhost:7777/api/objects/cluster");
    req.send();
};
