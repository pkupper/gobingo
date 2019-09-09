
let socket = new WebSocket("ws://localhost:8080/connect");
console.log("Attempting Connection...");

socket.onopen = () => {
    console.log("Successfully Connected");
    socket.send(JSON.stringify({type: 1, message: {username: "max", password: "secret"}}));
};

socket.onclose = event => {
    console.log("Socket Closed Connection: ", event);
};

socket.onerror = error => {
    console.log("Socket Error: ", error);
};

socket.onmessage = message => {
    console.log("Message: ", message);
};