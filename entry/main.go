package hello

import (
	"Controllers"
	"github.com/astaxie/beegae"
	"net/http"
)

func init() {
	beegae.Router("/", &Controllers.MainController{})
	beegae.Router("/user", &Controllers.UserController{})
	beegae.Router("/:id:int", &Controllers.UserController{})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		beegae.BeeApp.Handlers.ServeHTTP(w, r)
	})
}
