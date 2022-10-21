package main

import (
	"html/template"
	"net/http"
)

var templ *template.Template

func init() {
	templ = template.Must(template.ParseGlob("template/*.html"))

}

func home(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "index.html", nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "about.html", nil)
}

func signup(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "sign-up.html", nil)
}

func help(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "help.html", nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "contact.html", nil)
}
func main() {
	logo := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", logo))
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/help", help)
	http.HandleFunc("/contact", contact)

	http.ListenAndServe(":9999", nil)
}
