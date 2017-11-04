package routers

import (
	"github.com/astaxie/beego"
	"github.com/neatLines/beeDemo/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
}
