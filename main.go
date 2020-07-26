package main

import (
	"net/http"

	"github.com/jonparker/golang/controllers"
)

func main() {

	controllers.RegisterControllers()

	http.ListenAndServe(":3000", nil)
}
