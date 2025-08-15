let username = prompt("Enter your name:");
let socket;

document.addEventListener("DOMContentLoaded", function () {
  socket = new WebSocket("ws://" + window.location.host + "/ws");

  socket.onmessage = function (event) {
    const msg = JSON.parse(event.data);
    const li = document.createElement("li");

    // Distinguish between self and others
    if (msg.user === username) {
      li.classList.add("self");
    } else {
      li.classList.add("other");
    }

    li.innerHTML = `<strong>${msg.user}:</strong> ${msg.text}`;
    document.getElementById("messages").appendChild(li);
    scrollToBottom();
  };

  window.sendMessage = function () {
    let input = document.getElementById("message");
    let message = input.value.trim();
    if (message) {
      socket.send(JSON.stringify({ user: username, text: message }));
      input.value = '';
    }
  };
});

function scrollToBottom() {
  const messages = document.getElementById("messages");
  messages.scrollTop = messages.scrollHeight;
}
