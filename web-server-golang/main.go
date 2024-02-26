package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "The Path is wrong", http.StatusBadRequest)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Wrong Http Method", http.StatusNotAcceptable)
	}
	fmt.Fprintf(w, "Hello World")
}
func formFunc(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "error parsing the form")
		return
	}
	name := r.FormValue("name")
	fmt.Fprintf(w, name)
}
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formFunc)
	http.HandleFunc("/hello", helloFunc)
	fmt.Println("server starting on port :5000")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("err", err)
	}
}
