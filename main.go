package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
)

var tpl = template.Must(template.ParseFiles("index.html"))

// строки - срез байтов
func indexHandler(w http.ResponseWriter, r *http.Request) { // Обработчик для HTTP-запросов на путь "/"
	tpl.Execute(w, nil)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	params := u.Query()
	searchKey := params.Get("q")
	page := params.Get("p")
	if page == "" {
		page = "1"
	}

	fmt.Println("запрос ", searchKey)
	fmt.Println("страница ", page)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	//ServeMux - сопостовление url  с обработчиком запросов
	mux := http.NewServeMux() // новый маршрутизатор

	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/search", searchHandler)
	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
