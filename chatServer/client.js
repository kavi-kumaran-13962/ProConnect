var recipientId = document.getElementById("recipientId");
var messageElem = document.getElementById("message");
var output = document.getElementById("output");

// Function to generate a random UUID
function uuidv4() {
    return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
      var r = Math.random() * 16 | 0, v = c == 'x' ? r : (r & 0x3 | 0x8);
      return v.toString(16);
    });
}
// Generate a unique UUID for this client
var clientID = uuidv4();
var socket = new WebSocket(`ws://localhost:8082/ws?client-id=${clientID}`);

socket.onopen = function () {
    output.innerHTML += "Status: Connected\n";
};

socket.onclose = function(){
    output.innerHTML += "Status: Disconnected\n";
};

socket.onmessage = function (e) {
    const data = JSON.parse(e.data);
    output.innerHTML += "\nServer: " + data.message + "\n";
};

function sendMessage() {
    let recipientIdValue = recipientId.value;
    console.log(recipientIdValue)
    message = messageElem.value;
    socket.send(
      JSON.stringify({
        "to":recipientIdValue,
        message,
      })
    );
};