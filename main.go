package main

import (
	"fmt"
	"log"
	"net/http"
	"url-short/url"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/shorten", url.ShortenURLEndpoint).Methods("POST")
	router.HandleFunc("/shorten/{url}", url.GetOriginalUrlEndpoint).Methods("GET")
	router.HandleFunc("/shorten/{url}", url.UpdateLongURLEndpoint).Methods("PUT")
	router.HandleFunc("/shorten/{url}", url.DeleteURLEndpoint).Methods("DELETE")
	router.HandleFunc("/shorten/{url}/stats", url.GetStatsEndpoint).Methods("GET")
	fmt.Println("Server Started on Port :8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
