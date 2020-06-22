package controller

import (
	"net/http"
	"text/template"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("mysession"))

func Index(response http.ResponseWriter, request *http.Request) {
	tmp, _ := template.ParseFiles("views/index.html")
	tmp.Execute(response, nil)
}

func Login(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	email := request.Form.Get("email")
	//password := request.Form.Get("password")

	session, _ := store.Get(request, "mysession")
	session.Values["email"] = email
	session.Save(request, response)
	http.Redirect(response, request, "/welcome", http.StatusSeeOther)
}

func Logout(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	session.Options.MaxAge = -1
	session.Save(request, response)
	http.Redirect(response, request, "/index", http.StatusSeeOther)
}

func Welcome(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	email := session.Values["email"]
	data := map[string]interface{}{
		"email": email,
	}
	tmp, _ := template.ParseFiles("views/welcome.html")
	tmp.Execute(response, data)
}
