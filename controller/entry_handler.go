package controller

import (
	"net/http"
)

func RegisterControllers() {
	pc := newPostbackController()

	http.Handle("/postback", *pc)
	http.Handle("/postback/", *pc)
}
