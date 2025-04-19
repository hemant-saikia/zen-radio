package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	fmt.Println("Welcome to Zen radio!")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})
	http.HandleFunc("/radio", radio)

	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Printf("Server crash. Err [%v]", err.Error())
	}
}

func radio(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err = tmpl.Execute(w, map[string]string{
		"Title":    "Zen Radio",
		"VideoURL": "/static/video/bg-video-1.mp4",
		"AudioURL": "https://ice3.somafm.com/groovesalad-128-mp3", // or a stream URL
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
