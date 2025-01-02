package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var port = ":8000"
var movies []Movie

func main() {

	// creating router
	r := mux.NewRouter()

	movies = moviesList()

	// handler functions
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")

	// server setup
	fmt.Printf("server started on port: %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func moviesList() []Movie {

	movies = append(movies, Movie{ID: "1", ISBN: "1111", Title: "Jawan", Director: &Director{FirstName: "Shahrukh", LastName: "Khan"}})
	movies = append(movies, Movie{ID: "2", ISBN: "2222", Title: "Pathan", Director: &Director{FirstName: "Shahrukh", LastName: "Khan"}})

	return movies
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(movies)
	if err != nil {
		return
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000))

	err := json.NewEncoder(w).Encode(movie)
	if err != nil {
		return
	}

	movies = append(movies, movie)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	msg := map[string]string{"msg": "invalid id provided"}
	found := false

	for _, item := range movies {
		if item.ID == id {
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				return
			}
			found = true
			break
		}
	}

	if !found {
		err := json.NewEncoder(w).Encode(msg)
		if err != nil {
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	for index, item := range movies {
		if item.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	err := json.NewEncoder(w).Encode(movies)
	if err != nil {
		return
	}

	log.Fatalf("movie not found against given id: %s", id)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var updatedMovie Movie

	params := mux.Vars(r)
	id := params["id"]

	for index, item := range movies {

		if item.ID == id {
			movies[index] = Movie{
				ID:       id,
				Title:    updatedMovie.Title,
				ISBN:     updatedMovie.ISBN,
				Director: updatedMovie.Director,
			}

			err := json.NewEncoder(w).Encode(movies[index])

			if err != nil {
				return
			}
			return
		}
	}

	log.Fatalf("movie not found against given id: %s", id)
}
