package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(13) + 3
	y := rand.Intn(13) + 3
	index := make([][]int, y)
	for i := range index {
		index[i] = make([]int, x)
		for j := range index[i] {
			index[i][j] = rand.Intn(100) + 1
		}
	}
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, index)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	static := http.FileServer(http.Dir("src"))
	http.Handle("/src/", http.StripPrefix("/src/", static))
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
