package controllers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strings"
	"tcpserver/models"
)

var ws *websocket.Conn
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//Initwebsocket initialize web socket
func Initwebsocket(c *gin.Context) {
	log.Println("websocket initial")
	var err error
	ws, err = upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Can not setup websocket connection", err)
		return
	}
	/*if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(c.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		log.Println("Cannot setup WebSocket connection:", err)
		return
	}*/
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			log.Println("websoket is error")
			return
		}
		log.Println("receive webscoket message is ", string(p))
		var msg models.RecMessage
		json.Unmarshal(p, &msg)
		//l := logs.GetLogger()
		log.Println("The recMessage is ", msg.Name)
		handlerMessage(msg)
	}
}

func handlerMessage(msg models.RecMessage) {
	switch msg.Protocol {
	case "TCP":
		if strings.EqualFold(msg.Action, "Open") {
			log.Println("create server")
			CreateTCPServer(msg.Port)
		} else {
			log.Println("close server")
			CloseTCPServer()
		}

	}
	if msg.SaveFlag == 1 {
		SaveFlag = true
	} else {
		SaveFlag = false
	}

}

func sendWebSocket(data models.DataEvent) {
	senddata, err := json.Marshal(data)
	if err != nil {
		log.Println("Fail to marshal data", err)
	}
	ws.WriteMessage(websocket.TextMessage, senddata)
}
