package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["displayType"] = "TCP"
	c.Data["disableItem"] = true
	c.TplName = "index.tpl"
}
