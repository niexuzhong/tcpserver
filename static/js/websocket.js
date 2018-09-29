var socket;
var socketCreatedFlag=0;
var connectState=0;
var packageNumber=0;
var curSel=0;
$(document).ready(function () {
  //  CreateSocket();
  //  socketCreatedFlag=1;
});
function ChangeServerType() {
  var allString=['UDPButtonid','TCPButtonid','MQTTButonid','CoapButtonid'];
  var nameString=[' UDP',' TCP',' MQTT',' Coap'];
  curSel++;
  if(curSel>3)curSel=0;
  for(i=0;i<allString.length;i++)
  {
    if(i==curSel) {
       document.getElementById(allString[i]).className="label label-success bg-success ";
       document.getElementById("testServerName").innerHTML="Test "+nameString[curSel]+' Server'
     }
    else
      document.getElementById(allString[i]).className +="label-default bg-light disabled";
  }

}
function clickConnect() {
   var btnName=document.getElementById("connectBtnId");
      console.log("press connect button");
   if (connectState==0) {
       connectState=1;
       btnName.innerHTML="Disconnect";
       CreateSocket();
       if(socketCreatedFlag==1){
          sendOpenMethod(curSel);
       }
     } else {
     connectState=0;
     btnName.innerHTML="connect";
   }
 }
function addRow(inputdata) {
  var tab=document.getElementById("datatable");
  var tr=creatRow(inputdata);
  tab.appendChild(tr);
}
function creatRow(inputtxt) {
  var td,tr;
  tr=document.createElement("tr");
  for(var i=0;i<4;i++){
    td=document.createElement("td");

    td.appendChild(document.createTextNode(inputtxt[i]));
    tr.appendChild(td);
  }

  return tr;

}
function clickAddRow() {
  var indata=new Array(4);
  indata[0]=packageNumber;
  indata[1]="192.168.199.240";
  indata[2]="01 02 03";
  indata[3]="31 32 33";
  packageNumber++;
  addRow(indata);
}

  function sendOpenMethod(protocolType){
    var method=new Object();
    var protocolString = new Array('UDP','TCP','MQTT','COAP');
    method.name="sendCommand";
    method.port=9016;
    method.action="open";
    method.protocol=protocolString[protocolType];
    var methodstr=JSON.stringify(method);
    socket.send(methodstr);
  }
  function onOpen(evt){
     console.log("open data is "+evt.data);
  }
  function onReceive(evt) {
    var a=new Array(4);
    console.log("data receive is "+evt.data);
    var msg=JSON.parse(event.data);
    switch (msg.type){
      case "packageNumber":
             a[0]=msg.packageNumber;
             breka;
      case "address":
          a[0]=msg.address;
          break;
      case "hex":
          a[1]=msg.hex;
          break;
      case "ascii":
          a[3]=msg.ascii;
          break;
    }
    addRow(a);

  }
  function onClose(evt) {
    console.log("close data is "+evt.data);
  }
  function clickFileMethod(item) {
    CreateSocket()

  }
function CreateSocket() {
  socket=new WebSocket("ws://127.0.0.1:9001/ws/init");
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
