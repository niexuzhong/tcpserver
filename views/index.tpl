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
    <script src="static/js/jquery_3.2.1_jquery.slim.min.js" ></script>

    <script src="static/js/popper.js_1.12.9_umd_popper.min.js"></script>
    <script src="static/js/bootstrap.min.js"></script>
    <script src="static/js/websocket.js"></script>

</head>

<body>
  <nav class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 navbar-static-top">
      <a class="navbar-brand col-sm-3 col-md-2 mr-0" href="#" id="testServerName" onclick="ChangeServerType();">Test UDP Server</a>
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
              <!--  <a class="nav-link " id="UDPItemId" href="/UDP"> -->
                <!--  <span data-feather="file"></span>  -->
                  <h2  class="label label-success bg-success " id="UDPButtonid" name="button">UDP </h2>

              <!--  </a>  -->
              </li>
              <li class="nav-item ">
              <!--  <a class="nav-link " href="#"> -->
                <!--  <span data-feather="shopping-cart"></span>  -->
                  <h2  class="label label-default bg-light disabled" id="TCPButtonid" name="button">TCP </h2>

              <!--  </a>  -->
              </li>
              <li class="nav-item ">
              <!--  <a class="nav-link" href="#" >  -->
                <!--  <span data-feather="users" ></span>  -->
                  <h2 class="label label-default bg-light disabled" id="MQTTButonid" name="button">MQTT</h2>
              <!--  </a> -->
              </li>
              <li class="nav-item">
              <!--  <a class="nav-link" href="#">  -->
                <!--  <span data-feather="bar-chart-2"></span>  -->
                  <h2 class="label label-default bg-light disabled" id="CoapButtonid" name="button">COAP</h2>
            <!--    </a>  -->
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
                  <th>TimeStamp</th>
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



</body>

</html>
