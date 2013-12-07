package Controllers

import (
	"github.com/astaxie/beegae"
)

type MainController struct {
	beegae.Controller
}

func (this *MainController) Get() {
	this.Ctx.Output.Body([]byte("hello"))
}
