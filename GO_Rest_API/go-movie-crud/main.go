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

type Movie struct {
	ID       string    `json:"id"`
	Isdn     string    `json:"isdn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if string(item.ID) == (params["id"]) {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // getting prams from the user
	var curr_movie Movie
	_ = json.NewDecoder(r.Body).Decode(&curr_movie)
	for index, item := range movies {
		if string(item.ID) == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			curr_movie.ID = params["id"]
			movies = append(movies, curr_movie)
			json.NewEncoder(w).Encode(curr_movie)
			return
		}
	}
}

// START HERE https://www.youtube.com/watch?v=TkbhQQS3m_o

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var curr_movie Movie
	_ = json.NewDecoder(r.Body).Decode(&curr_movie)
	curr_movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, curr_movie)
	json.NewEncoder(w).Encode(curr_movie)

}
func getMovieByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // getting prams from the user
	for _, item := range movies {
		if string(item.ID) == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "Movie not found"}`))
}
func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{
		ID:       "1",
		Isdn:     "123456",
		Title:    "Movie One",
		Director: &Director{FirstName: "John", LastName: "Doe"},
	})
	movies = append(movies, Movie{
		ID:       "2",
		Isdn:     "123457",
		Title:    "Movie Two",
		Director: &Director{FirstName: "Steve", LastName: "Smith"},
	})
	r.HandleFunc("/", testServer).Methods("GET")
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovieByID).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Println("Server on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func testServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Server is running")
}
