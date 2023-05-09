package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1><p><a href=\"http://localhost:3000/contact\">Contact page</a>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Us</h1><p>To get in touch, email at <a href=\"mailto:jeppe798@hotmail.com\">jeppe798@hotmail.com</a><p><a href=\"http://localhost:3000\">Back to home</a>")
}

//

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "Sorry - the page you were looking for could not be found :(", http.StatusNotFound)
	}
}

func main() {
	var router Router
	fmt.Println("Starting server on :3000....")
	http.ListenAndServe(":3000", router)

}
