package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

type Home struct {
}
type Recipie struct {
}

func (home Home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Recipies HomePage")
}
func (recipie Recipie) CreateRecipie(w http.ResponseWriter, r *http.Request) {

}
func (recipie Recipie) DeleteRecipie(w http.ResponseWriter, r *http.Request) {

}
func (recipie Recipie) AddRecipie(w http.ResponseWriter, r *http.Request) {

}
func (recipie Recipie) UpdateRecipie(w http.ResponseWriter, r *http.Request) {

}
func (recipie Recipie) getRecipie(w http.ResponseWriter, r *http.Request) {

}
func (recipie Recipie) allRecipie(w http.ResponseWriter, r *http.Request) {

}

var (
	RecipeRe       = regexp.MustCompile(`^/recipes/*$`)
	RecipeReWithID = regexp.MustCompile(`^/recipes/([a-z0-9]+(?:-[a-z0-9]+)+)$`)
)

func (recipie Recipie) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost && RecipeRe.MatchString(r.URL.Path):
		recipie.CreateRecipie(w, r)
		return
	case r.Method == http.MethodGet && RecipeReWithID.MatchString(r.URL.Path):
		recipie.getRecipie(w, r)
		return
	case r.Method == http.MethodGet && RecipeRe.MatchString(r.URL.Path):
		recipie.allRecipie(w, r)
		return
	case r.Method == http.MethodPut && RecipeReWithID.MatchString(r.URL.Path):
		recipie.UpdateRecipie(w, r)
		return
	case r.Method == http.MethodDelete && RecipeReWithID.MatchString(r.URL.Path):
		recipie.DeleteRecipie(w, r)
		return

	}
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
