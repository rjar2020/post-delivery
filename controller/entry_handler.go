package controller

import (
	"net/http"
)

//RegisterControllers is the entrypoint of the postback API
func RegisterControllers() {
	pc := newPostbackController()

	http.Handle("/postback", *pc)
	http.Handle("/postback/", *pc)
}
