package main

import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie{
	ID string 'json: "id"'
	Isbn string 'json: "isbn"'
	Title string 'json: "title"'
	Director *Director 'json: "director"'
}

type Director{
	Name string 'json: "name"'
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies
}

func deleteMovie(w http.ResponseWriter, r *http.Request){

}

func deleteMovie(w http.ResponseWriter, r *http.Request){

}

func deleteMovie(w http.ResponseWriter, r *http.Request){

}
func main(){
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isb: "123", Title: "Movie 1", Director: &Director{Name:"Spielberg"}})
	movies = append(movies, Movie{ID: "2", Isb: "1234", Title: "Movie 2", Director: &Director{Name:"Nolan"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{i}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{i}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{i}", deleteMovie).Methods("Delete")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))
}