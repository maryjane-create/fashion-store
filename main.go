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

func main() {
	logo := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", logo))
	http.HandleFunc("/", home)
	http.ListenAndServe(":9999", nil)
}
