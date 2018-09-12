package controllers

import (
	"fmt"
	"net"
	"strconv"
	"tcpserver/models"
	"time"

	"github.com/astaxie/beego"
)

var Echoflag int

func CreateTCPServer(port int, IPAddress string) error {

	tcpListener, err := net.Listen(IPAddress, ":"+strconv.Itoa(port))
	if err != nil {
		beego.Error("create tcp server error")
		beego.Error(err.Error())
		return err
	}
	defer tcpListener.Close()
	for {
		conn, err := tcpListener.Accept()
		if err != nil {
			beego.Error("accept ecp server error")
			return err
		}
		go handleRequest(conn)
	}

	return nil
}
func SendtoWebSocket(senddata []byte) {
	var data models.DataEvent
	data.Timestamp = time.Now().String()
	data.Asciistring = fmt.Sprintf("%s", senddata)
	data.Hexstring = fmt.Sprintf("%x", senddata)
	sendWebSocket(data)
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
		SendtoWebSocket(buffer)

	}
	conn.Close()
}
