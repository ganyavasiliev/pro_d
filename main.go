package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("index.html")
	tpl.Execute(w, nil)
}
func main() {
	http.HandleFunc("/", indexHandler)
	fs_css := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs_css))
	fs_image := http.FileServer(http.Dir("image"))
	http.Handle("/image/", http.StripPrefix("/image/", fs_image))
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
