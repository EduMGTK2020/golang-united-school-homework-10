package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{PARAM}", handleName).Methods("GET")
	router.HandleFunc("/bad", handleBad).Methods("GET")
	router.HandleFunc("/data", handleData).Methods("POST")
	router.HandleFunc("/headers", handleHeaders).Methods("POST")

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

func handleName(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["PARAM"]
	responce := fmt.Sprintf("Hello, %v!", param)
	w.Write([]byte(responce))
}
func handleBad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}
func handleData(w http.ResponseWriter, r *http.Request) {
	param, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var b bytes.Buffer
	b.WriteString("I got message:\n")
	b.Write(param)
	w.Write(b.Bytes())

}
func handleHeaders(w http.ResponseWriter, r *http.Request) {
	fieldA, err := strconv.Atoi(r.Header.Get("a"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fieldB, err := strconv.Atoi(r.Header.Get("b"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("a+b", fmt.Sprint(fieldA+fieldB))
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
