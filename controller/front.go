package controller

import "net/http"

func RegisterControllers() {

	uc := newUserController()
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc) // not the same /test and /test/

}
