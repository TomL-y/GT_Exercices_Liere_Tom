package main

import (
	"net/http"
	"text/template"
)

func errorpage(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		tmpl := template.Must(template.ParseFiles("error.html"))
		tmpl.Execute(w, nil)
	}
}

func main() {
	static := http.FileServer(http.Dir("src"))
	http.Handle("/src/", http.StripPrefix("/src/", static))
	tmpl1 := template.Must(template.ParseFiles("index1.html"))
	tmpl2 := template.Must(template.ParseFiles("index2.html"))
	tmpl3 := template.Must(template.ParseFiles("index3.html"))
	http.HandleFunc("/index1", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/index1" {
			errorpage(w, r, http.StatusNotFound)
			return
		} else {
			tmpl1.Execute(w, nil)
		}
	})
	http.HandleFunc("/index2", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/index2" {
			errorpage(w, r, http.StatusNotFound)
			return
		} else {
			tmpl2.Execute(w, nil)
		}
	})
	http.HandleFunc("/index3", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/index3" {
			errorpage(w, r, http.StatusNotFound)
			return
		} else {
			tmpl3.Execute(w, nil)
		}
	})
	http.ListenAndServe(":8080", nil)
}
