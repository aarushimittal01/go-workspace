package main

import (
	"net/http"

	controller "github.com/aarushimittal01/social_login/controllers"
)

func main() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/index", controller.Index)
	http.HandleFunc("/login", controller.Index)
	http.HandleFunc("/welcome", controller.Welcome)
	http.HandleFunc("/logout", controller.Logout)
	http.ListenAndServe(":3000", nil)
}
