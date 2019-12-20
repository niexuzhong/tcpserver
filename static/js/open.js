var table =new Tabulator("#example-table",{
	headerSort:false,
	selectable:false,
	columns:[
	  {title:"TimeStamp",field:"timestamp"},
	  {title:"IPAddress",field:"ipaddress"},
	  {title:"Hex",field:"hex"},
	  {title:"ASCII",field:'ascii'},
	],
});

function createRow(inputtxt) {
  data=new Object();
  data.timestamp=inputtxt[0];
  data.ipaddress=inputtxt[1];
  data.hex=inputtxt[2];
  data.ascii=inputtxt[3];
  table.addRow(data);
}