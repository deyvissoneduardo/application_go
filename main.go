package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func connectionDatabase() *sql.DB {
	connection := "user=admin dbname=store password=admin host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Product struct {
	Name        string
	Description string
	Price       float64
	Amount      int
}

func main() {
	db := connectionDatabase()
	defer db.Close()
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
