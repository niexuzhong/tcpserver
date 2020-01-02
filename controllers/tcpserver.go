package controllers

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"tcpserver/models"
	"time"

	"github.com/gin-gonic/gin"
)

//Echoflag echo or not
var EchoFlag bool
var remoteAddr net.Addr
var server net.Listener
var conn net.Conn
var packageNumber int
var socketOpenFlag = false

//SaveFlag save flag
var SaveFlag bool

//IndexHandler handler the index.html
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"displayType": "TCP",
	})
}

//CreateTCPServer create TCP server
func CreateTCPServer(port int) error {
	var err error
	if socketOpenFlag == true {
		return nil
	}
	server, err = net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Println("create tcp server error")
		log.Println(err.Error())
		return err
	}
	socketOpenFlag = true
	//log.Println("port is", strconv.Itoa(port))
	//log.Println("begin to listen")
	//defer tcpListener.Close()
	go serverTask(server)

	return nil
}

//CloseTCPServer close tcp server
func CloseTCPServer() {
	socketOpenFlag = false
	fmt.Println("close tcp server")
	if conn != nil {
		conn.Close()
	}
	if server != nil {
		server.Close()
	}

}

func serverTask(listener net.Listener) error {
	var err error
	log.Println("begin server task")
	for {
		conn, err = listener.Accept()
		if err != nil {

			log.Println("accept tcp server error")
			log.Println(err.Error())
			return err
		}
		remoteAddr = conn.RemoteAddr()
		log.Println("the remote address is", remoteAddr)
		go handleRequest(conn)
		if SaveFlag == true {
			log.Println("create save data task ")
			packageNumber = 0
			models.InitSaveChan()
			go models.CreateDataSaveTask("../datafile/"+remoteAddr.String())
		}

	}

}

func sendtoWebSocket(senddata []byte) {
	var data models.DataEvent
	timeTemplate1 := "2006-01-02 15:04:05"
	data.PackageNumber = packageNumber
	t := (int64)(time.Now().Unix())
	data.TimeStamp = time.Unix(t, 0).Format(timeTemplate1)
	//data.TimeStamp = strconv.FormatInt(time.Now().Unix(), 10)
	data.ASCIIString = fmt.Sprintf("%s", senddata)
	data.Address = fmt.Sprintf("%s", remoteAddr)
	data.HexString = fmt.Sprintf("%x", senddata)
	sendWebSocket(data)
	if SaveFlag == true {
		log.Println("transimit save data")
		models.TranSaveChan(data)
	}
	packageNumber++

}

func handleRequest(conn net.Conn) {

	for {
		buffer := make([]byte, 1024)
		reqLen, err := conn.Read(buffer)
		if err != nil {
			log.Println("error is ", err.Error())
			break
		}
		log.Println("recLen is", reqLen)

		buffer = buffer[:reqLen]
		if EchoFlag != false {
			conn.Write(buffer)
		}
		hexstring := fmt.Sprintf("%02x ", buffer)
		log.Println("hexString is ", hexstring)
		if getWebsocketStatus() == true {
			sendtoWebSocket(buffer)
		}

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

func SetEchoFlag(enable bool) {
	EchoFlag = enable
}