package routers

import (
	"chaincodeUI/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/chaincode/", &controllers.ChaincodeController{})
}
