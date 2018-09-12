package main

import (
	_ "tcpserver/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

