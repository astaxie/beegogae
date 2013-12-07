package hello

import (
	"Controllers"
	"github.com/astaxie/beegae"
	"github.com/astaxie/beegae/middleware"
	"net/http"
)

func init() {
	beegae.Router("/", &Controllers.MainController{})
	beegae.Router("/user", &Controllers.UserController{})
	beegae.Router("/auth/login", &Controllers.UserController{})
	beegae.Router("/:id:int", &Controllers.UserController{})

	err := beegae.BuildTemplate(beegae.ViewsPath)
	if err != nil {
		panic(err)
	}

	middleware.VERSION = beegae.VERSION
	middleware.AppName = beegae.AppName
	middleware.RegisterErrorHander()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		beegae.BeeApp.Handlers.ServeHTTP(w, r)
	})
}
