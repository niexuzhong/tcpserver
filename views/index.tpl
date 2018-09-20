<!DOCTYPE html>
<html lang="en">

<head>

    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <meta name="description" content="">
    <meta name="author" content="">
  <!--  <link rel="icon" href="../../../../favicon.ico">  -->

    <title>Test Server</title>

    <!-- Bootstrap core CSS -->
    <link href="static/css/bootstrap.min.css" rel="stylesheet">

    <!-- Custom styles for this template -->
    <link href="static/css/dashboard.css" rel="stylesheet">


</head>

<body>
  <nav class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 navbar-static-top">
      <a class="navbar-brand col-sm-3 col-md-2 mr-0" href="#" onclick="ChangeServerType();">Test Server</a>
    <!--  <form class="form-inline col-md-6 input-group" >  -->
      <!--  <label class="form-control form-control-dark col-md-2 ">Port Number </label>
        <input type="text" class="form-control form-control-light col-md-2">
      </form> -->
      <div class="input-group ">
        <div class="input-group-prepend">
          <div class="input-group-text">Log File
            <input type="checkbox" aria-label="Checkbox for following text input">
          </div>
        </div>
        <input type="text" class="form-control" aria-label="Text input with checkbox">
      </div>
      <!--<input class="form-control form-control-dark w-100" type="text" placeholder="Search" aria-label="Search">-->
    <!--  <ul class="navbar-nav px-3">
        <li class="nav-item text-nowrap">
          <a class="nav-link" href="#">Sign out</a>
        </li>
      </ul>  -->
      <span class="col-md-6"> </span>
       <button class="btn btn-sm btn-outline-secondary p-3 " onclick="clickConnect();" id="connectBtnId">Connect</button>
    </nav>

   <!-- <div class="container-fluid">   -->
      <div class="row">
        <nav class="col-md-2 d-none d-md-block bg-dark sidebar">
          <div class="sidebar-sticky">
            <ul class="nav flex-column">

              <li class="nav-item ">
                <a class="nav-link " id="UDPItemId" href="/UDP">
                <!--  <span data-feather="file"></span>  -->
                  <button type="button" class="btn btn-success btn-lg " id="UDPButtonid" name="button">UDP </button>

                </a>
              </li>
              <li class="nav-item ">
                <a class="nav-link " href="#">
                <!--  <span data-feather="shopping-cart"></span>  -->
                  <button type="button" class="btn btn-success btn-lg disabled" id="TCPButtonid" name="button">TCP </button>

                </a>
              </li>
              <li class="nav-item ">
                <a class="nav-link" href="#" >
                <!--  <span data-feather="users" ></span>  -->
                  <button type="button" class="btn btn-success btn-lg disabled" id="MQTTButonid" name="button">MQTT</button>
                </a>
              </li>
              <li class="nav-item">
                <a class="nav-link" href="#">
                <!--  <span data-feather="bar-chart-2"></span>  -->
                  <button type="button" class="btn btn-success btn-lg disabled" id="CoapButtonid" name="button">COAP</button>
                </a>
              </li>
            </ul>
          </div>
        </nav>

        <main role="main" class="col-md-9 ml-sm-auto col-lg-10 pt-3 px-4">
          <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pb-2 mb-3 border-bottom">
            <h1 class="h2">{{.displayType}}</h1>
            <div class="btn-toolbar mb-2 mb-md-0">
              <div class="btn-group mr-2">
                <button class="btn btn-sm btn-outline-secondary" onclick="clickAddRow();">Download</button>

              </div>

            </div>
          </div>

          <h2>Data table</h2>
          <div class="table-responsive">
            <table class="table table-striped table-sm" id="datatable">
              <thead>
                <tr>
                  <th>#</th>
                  <th>IPAddress</th>
                  <th>Hex</th>
                  <th>ASCII</th>
                </tr>
              </thead>
              <tbody>


              </tbody>
            </table>
          </div>
        </main>
      </div>
    <!-- </div>  -->
    <script type="text/javascript">
      var connectState=0;
      var packageNumber=0;
      var curSel=0;
      function ChangeServerType() {
        var allString=['UDPButtonid','TCPButtonid','MQTTButonid','CoapButtonid'];

        for(i=0;i<allString.length;i++)
        {
          if(i==curSel)
             document.getElementById(allString[i]).className="btn btn-success btn-lg";
          else
            document.getElementById(allString[i]).className +="btn btn-success btn-lg disabled";
        }
        curSel++;
        if(curSel>3)curSel=0;
      }
      function clickConnect() {
        var disableitem=document.getElementById("UDPItemId");
            disableitem.className+="disabled";
        var btnName=document.getElementById("connectBtnId");
         if (connectState==0) {
           connectState=1;
           btnName.innerHTML="Disconnect";

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
        function sendOpenMethod(){
          var method=new object();
          method.name="sendCommand";
          method.port=9016;
          method.action="open";
          method.protocol="UDP";
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
    </script>
    <script src="static/js/jquery_3.2.1_jquery.slim.min.js" ></script>

    <script src="static/js/popper.js_1.12.9_umd_popper.min.js"></script>
    <script src="static/js/bootstrap.min.js"></script>

</body>

</html>
