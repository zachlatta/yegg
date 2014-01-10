package main

import (
	"fmt"
	"net/http"
)

const (
	USER = "jdoe"
	PASS = "foobar"
)

func main() {
	http.HandleFunc("/user_login_submit", loginHandler)
	fmt.Println("Server started successfully.")
	http.ListenAndServe(":1759", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	username, password := r.FormValue("userName"), r.FormValue("x")
	if username == USER && password == PASS {
		fmt.Fprintln(w, "<h1>Login Succeeded</h1>")
	} else {
		fmt.Fprintln(w, "<h1>Login Failed</h1>")
	}
}
