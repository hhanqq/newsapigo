package main

import (
	"html/template"
	"net/http"
	"os"
)

var tpl = template.Must(template.ParseFiles("index.html"))

// строки - срез байтов
func indexHandler(w http.ResponseWriter, r *http.Request) { // Обработчик для HTTP-запросов на путь "/"
	tpl.Execute(w, nil)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	//ServeMux - сопостовление url  с обработчиком запросов
	mux := http.NewServeMux() // новый маршрутизатор

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
