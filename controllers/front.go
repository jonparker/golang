package controllers

import (
	"net/http"
)

//RegisterControllers for routes
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
