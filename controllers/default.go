package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	type Tmp struct {
		HerfURL  string
		HerfName string
	}
	collection:=[]Tmp{}
	localhost:="localhost:8080/chaincode?operation="
	
	collection=append(collection,Tmp{localhost+SVN,"更新代码"})
	collection=append(collection,Tmp{localhost+INIT,"部署合约"})
	collection=append(collection,Tmp{localhost+UPDATE,"更新合约"})
	collection=append(collection,Tmp{localhost+CLEAR,"清除合约"})
	
	c.Data["Colletion"] = collection
	c.TplName = "index.tpl"
}
