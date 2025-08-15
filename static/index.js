let socket = new WebSocket("ws://" + window.location.host + "/ws");

socket.onmessage = function (event) {
    console.log(event.data);
    document.querySelector('#messages');
    
    let li = document.createElement('li');
    li.textContent = event.data;
    document.querySelector('#messages').appendChild(li);
}

function sendMessage() {
    let input = document.getElementById("#message");
    socket.send(input.value)
    input.value = '';
}