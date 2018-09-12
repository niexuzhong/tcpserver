package routers

import (
	"tcpserver/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/ws/init", &controllers.WebsocketController{}, "get:Initwebsocket")
}
