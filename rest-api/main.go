package main

import (
	"fmt"
	"log"
	"net/http"
)

type Home struct {
}
type Recipie struct {
}

func (home Home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Recipies HomePage")
}
func (recipie Recipie) CreateRecipie(w http.ResponseWriter, r *http.Response) {

}
func (recipie Recipie) DeleteRecipie(w http.ResponseWriter, r *http.Response) {

}
func (recipie Recipie) AddRecipie(w http.ResponseWriter, r *http.Response) {

}
func (recipie Recipie) UpdateRecipie(w http.ResponseWriter, r *http.Response) {

}
func main() {
	mux := http.NewServeMux()
	mux.Handle("/", Home{})
	mux.Handle("/recipies", Recipie{})
	mux.Handle("/recipies", Recipie{})
	fmt.Printf("server running on port 5000")
	err := http.ListenAndServe(":5000", mux)
	if err != nil {
		log.Fatal("server error")
	}
}
