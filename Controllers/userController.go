package Controllers

import (
	"fmt"
	"github.com/astaxie/beegae"
)

type UserController struct {
	beegae.Controller
}

func (this *UserController) Get() {
	fmt.Println(this.Ctx.Input.Params)
	this.Ctx.Output.Body([]byte("i am user" + this.Ctx.Input.Param(":id")))
}
