package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.ParseFiles("index.html")
		if err != nil {
			log.Print(err)
			w.WriteHeader(501)
			return
		}
		err = temp.Execute(w, nil)
		if err != nil {
			log.Print(err)
			w.WriteHeader(501)
			return
		}
		w.WriteHeader(200)
	})
	mux.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.ParseFiles("index.html")
		if err != nil {
			log.Print(err)
			w.WriteHeader(501)
			return
		}
		err = temp.Execute(w, nil)
		if err != nil {
			log.Print(err)
			w.WriteHeader(501)
			return
		}
		w.WriteHeader(200)
	})
	panic(http.ListenAndServe(":9996", mux))

}
