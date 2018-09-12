var socket;
function CreateSocket() {
  socket=new WebSocket("ws://127.0.0.1/ws/init");
  socket.onopen=function(evt) {
    onOpen(evt);
  };
  socket.onmessage=function(evt) {
    onReceive(evt);
  };
  socket.onclose=function(evt) {
    onClose(evt);
  };
}
function onOpen(evt) {
  console.log('open data is '+evt.data);
}
function onReceive(evt) {
  var msg=JSON.parse(evt.data);
  console.log('timestamp is '+msg.Timestamp+'Hexstring is '+msg.Hexstring+'Asciistring is '+msg.Asciistring);

}
function onClose(evt) {
  console.log('close data is '+evt.data);
}
