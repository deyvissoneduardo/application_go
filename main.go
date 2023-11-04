package main

import (
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

type Product struct {
	Name        string
	Description string
	Price       float64
	Amount      int
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8181", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{"Iphone 15", "Apple", 6.500, 15},
		{"Iphone 15 Plus", "Apple", 6.050, 10},
		{"Iphone 15 Pro", "Apple", 7.030, 5},
	}
	temp.ExecuteTemplate(w, "Index", products)
}
