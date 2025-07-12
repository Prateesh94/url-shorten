package url

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type Req struct {
	URL string `json:"url"`
}
type UrlData struct {
	Id     int    `json:"id"`
	Url    string `json:"url"`
	Short  string `json:"shortCode"`
	Create string `json:"createAt"`
	Update string `json:"updateAt"`
	hits   int
}

var (
	mu sync.Mutex
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateShortURL() string {
	rand.NewSource(time.Now().UnixNano())
	b := make([]byte, 6)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
func ShortenURLEndpoint(w http.ResponseWriter, r *http.Request) {
	var req Req
	var short Req
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	mu.Lock()
up:
	var dt UrlData
	short.URL = generateShortURL()
	dt, er := addUrl(short.URL, req.URL)
	if er != nil {
		goto up
	}
	mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dt)
}

func GetOriginalUrlEndpoint(w http.ResponseWriter, r *http.Request) {
	var req Req
	var short UrlData
	d := mux.Vars(r)
	u := d["url"]
	if u == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "URL NOT FOUND")
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	mu.Lock()
	short, er := retrieveUrl(u)
	mu.Unlock()
	if er != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "URL NOT FOUND")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(short)
}

func UpdateLongURLEndpoint(w http.ResponseWriter, r *http.Request) {
	var req Req
	var short UrlData
	d := mux.Vars(r)
	u := d["url"]
	if u == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "URL NOT FOUND")
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	mu.Lock()
	short, er := updateurl(u, req.URL)
	mu.Unlock()
	if er != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "URL NOT FOUND")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(short)
}

func DeleteURLEndpoint(w http.ResponseWriter, r *http.Request) {

	d := mux.Vars(r)
	u := d["url"]
	if u == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "URL NOT FOUND")
		return
	}
	mu.Lock()
	er := deleteurl(u)
	mu.Unlock()
	if er != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "URL NOT FOUND", http.StatusNotFound)
		return
	} else {
		w.WriteHeader(http.StatusNoContent)
		fmt.Fprintln(w, "No Content", http.StatusNoContent)
	}
}

func GetStatsEndpoint(w http.ResponseWriter, r *http.Request) {
	d := mux.Vars(r)
	var short UrlData
	type Show struct {
		Id     int    `json:"id"`
		Url    string `json:"url"`
		Short  string `json:"shortCode"`
		Create string `json:"createAt"`
		Update string `json:"updateAt"`
		Hits   int    `json:"accessCount"`
	}
	u := d["url"]
	if u == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "URL NOT FOUND", http.StatusNotFound)
		return
	}
	short, er := retrieveUrl(u)
	p := Show{short.Id, short.Url, short.Short, short.Create, short.Update, short.hits}
	if er != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "URL NOT FOUND", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
