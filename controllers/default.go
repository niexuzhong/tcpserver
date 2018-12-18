package controllers

import (
	"github.com/astaxie/beego"
)

//MainController main controller
type MainController struct {
	beego.Controller
}

//Get get request
func (c *MainController) Get() {
	c.Data["displayType"] = "TCP"
	c.Data["disableItem"] = true
	c.TplName = "index.tpl"
}
