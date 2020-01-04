var socket;
var socketCreatedFlag=0;
var connectState=0;
var packageNumber=0;
var curSel=0;
var SaveFlag=0;
var EchoFlag=0;
var url = window.location.href;
url=GetUrl(url);
function clickUpLoad(e) {
  var fileInfo = e.files[0];
  var fileName = fileInfo.name;
  var filesize = fileInfo.size;
  console.log("name is "+fileName+"size is "+filesize);
  var fileReader = new FileReader();
  fileReader.readAsArrayBuffer(fileInfo);
  fileReader.onload = function (result){
    var pakoString = getUploadingFileContentString(this.result);
    var urlstring = "http://"+url+"/upload";
    console.log("ajax url is "+urlstring);
    $.ajax({
      url:urlstring,
      type:"POST",
      data:{
        fileContent:pakoString,
        fileName:fileName,
      },
      success:function(data) {
        if (data=="True") {
          alert("upload successful");
        } else {
          alert("upload failure");
        }
      },
      error: function(e){
        alert(e.responseText);
      }
    })
  }
}
function getUploadingFileContentString(readResult) {
  if(readResult == null ){
    return null
  }
    var fileContentArr = new Uint8Array(readResult);
  var fileContentStr = "";
  for (var i=0;i<fileContentArr.length;i++){
    fileContentStr+=String.fromCharCode(fileContentArr[i]);
  }
//  var pakoString=fromByteArray(fileContentStr);
  return fileContentStr;
}
$(document).ready(function () {
    CreateSocket();
    socketCreatedFlag=1;
});

function GetUrl( herf){
    var index= herf.indexOf("/",7);
     herf=herf.slice(7,index);
    console.log("url is",herf);
    return herf;
}
function ChangeServerType() {
  let allString=['UDPButtonid','TCPButtonid']; //,'MQTTButonid','CoapButtonid'];
  let nameString=[' UDP',' TCP'];  //MQTT',' Coap'];
  curSel++;
  //if(curSel>3)curSel=0;
  curSel=curSel%allString.length;
  console.log('cursel is ',curSel)
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

function clickEchoBtn() {
    if (document.getElementById("echoCheckbox").checked==true) {
        EchoFlag=1;
    }else {
        EchoFlag=0;
    }
}

function clickSaveBtn() {
  if (document.getElementById("saveCheckbox").checked==true) {
    SaveFlag=1;
  }else {
    SaveFlag=0;
  }
}
function clickConnect() {
   var btnName=document.getElementById("connectBtnId");
      console.log("press connect button");
   if (connectState==0) {
       connectState=1;
       btnName.innerHTML="Disconnect";
       if(socketCreatedFlag==1){
          sendOpenMethod(curSel,true);
       }
     } else {

     sendOpenMethod(curSel,false) ;
     connectState=0;
     btnName.innerHTML="connect";
   }
 }
function addRow(input_data) {
 createRow(input_data)
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

  function sendOpenMethod(protocolType,flag){
    var method=new Object();
    var protocolString = new Array('UDP','TCP','MQTT','COAP');
    method.name="sendCommand";
    method.port=9016;
    if(flag==true){
      method.action="open";
    }   else {
        method.action="close";
       }
    method.protocol=protocolString[protocolType];
    method.saveFlag=SaveFlag
    method.echoFlag=EchoFlag
    var methodstr=JSON.stringify(method);
    socket.send(methodstr);
  }
  function onOpen(evt){
     console.log("open data is "+evt.data);
  }
  function onReceive(evt) {
    var a=new Array(4);
    console.log("data receive is "+evt.data);
    var msg=JSON.parse(evt.data);
     a[0]=msg.TimeStamp;
     a[1]=msg.Address;
     a[2]=msg.HexString;
     a[3]=msg.ASCIIString;
     addRow(a);
  }
  function onClose(evt) {
    console.log("close data is "+evt.data);
  }
  function clickFileMethod(item) {
    CreateSocket()

  }
function CreateSocket() {

  wsurl="ws://"+url+"/ws/init"
  console.log("url is ",wsurl);
  socket=new WebSocket(wsurl);
  console.log("create websocket");
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
