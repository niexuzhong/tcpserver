package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tcpserver/models"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

type WebsocketController struct {
	beego.Controller
}

var ws *websocket.Conn

func (c *WebsocketController) Initwebsocket() {
	fmt.Println("websocket initial")
	var err error
	ws, err = websocket.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(c.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("websoket is error")
			return
		}
		fmt.Println("receive webscoket message is ", string(p))

	}
}

func sendWebSocket(data models.DataEvent) {
	senddata, err := json.Marshal(data)
	if err != nil {
		beego.Error("Fail to marshal data", err)
	}
	ws.WriteMessage(websocket.TextMessage, senddata)
}