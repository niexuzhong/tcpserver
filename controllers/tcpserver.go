package controllers

import (
	"fmt"
	"net"
	"strconv"
	"tcpserver/models"
	"time"

	"github.com/astaxie/beego"
)

//Echoflag echo or not
var Echoflag int
var remoteAddr net.Addr
var server net.Listener
var conn net.Conn
var packageNumber int

//SaveFlag save flag
var SaveFlag bool

//CreateTCPServer create TCP server
func CreateTCPServer(port int) error {
	var err error
	server, err = net.Listen("tcp4", "127.0.0.1:"+strconv.Itoa(port))
	if err != nil {
		beego.Error("create tcp server error")
		beego.Error(err.Error())
		return err
	}
	beego.Info("port is", port)
	//defer tcpListener.Close()
	go serverTask(server)

	return nil
}

//CloseTCPServer close tcp server
func CloseTCPServer() {
	fmt.Println("close tcp server")
	conn.Close()
	server.Close()
}

func serverTask(listener net.Listener) error {
	var err error

	for {
		conn, err = listener.Accept()
		if err != nil {

			beego.Error("accept tcp server error")
			beego.Error(err.Error())
			return err
		}
		remoteAddr = conn.RemoteAddr()
		beego.Info("the remote address is", remoteAddr)
		go handleRequest(conn)
		if SaveFlag == true {
			beego.Info("create save data task ")
			packageNumber = 0
			models.InitSaveChan()
			go models.CreateDataSaveTask(remoteAddr.String())
		}

	}

}

func sendtoWebSocket(senddata []byte) {
	var data models.DataEvent
	data.PackageNumber = packageNumber
	data.TimeStamp = strconv.FormatInt(time.Now().Unix(), 10)
	data.ASCIIString = fmt.Sprintf("%s", senddata)
	data.Address = fmt.Sprintf("%s", remoteAddr)
	data.HexString = fmt.Sprintf("%x", senddata)
	sendWebSocket(data)
	if SaveFlag == true {
		beego.Info("transimit save data")
		models.TranSaveChan(data)
	}
	packageNumber++

}

func handleRequest(conn net.Conn) {
	buffer := make([]byte, 1024)
	for {
		reqLen, err := conn.Read(buffer)
		if err != nil {
			beego.Error("read buffer error")
			break
		}

		if Echoflag != 0 {
			conn.Write(buffer[reqLen:])
		}
		buffer = buffer[:reqLen]
		sendtoWebSocket(buffer)

	}
	if SaveFlag == true {
		models.EndSaveTask()
		time.Sleep(time.Duration(400) * time.Millisecond)
	}

	conn.Close()
}

//SetSaveFlag set save flag
func SetSaveFlag(enable bool) {
	SaveFlag = enable
}
