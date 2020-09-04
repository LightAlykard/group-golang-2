package routers

import (
	"DZ5/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/messeges", &controllers.MessegeController{})
	beego.Router("/messege/:id([0-9]+)", &controllers.MessegeController{}, "get:GetOneMessege;put:UpdateMessege;delete:DeleteMessege")
}
