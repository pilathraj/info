package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type List struct {
	ID   int    `json:"Id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var items = []List{
	{
		ID:   1,
		Name: "Apple",
		Age:  26,
	},
}

func getLists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
func getList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, v := range items {
		if strconv.Itoa(v.ID) == params["id"] {
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	json.NewEncoder(w).Encode(&List{})
}
func deleteList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for k, v := range items {
		if strconv.Itoa(v.ID) == params["id"] {
			items = append(items[:k], items[k+1:]...)
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	json.NewEncoder(w).Encode(&List{})
}

func createList(w http.ResponseWriter, r *http.Request) {
	var item List
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		panic(err)
	}
	items = append(items, item)
	json.NewEncoder(w).Encode(item)
}

func updateList(w http.ResponseWriter, r *http.Request) {
	var tempitem List
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&tempitem)
	if err != nil {
		panic(err)
	}
	var params = mux.Vars(r)

	for k, v := range items {
		if strconv.Itoa(v.ID) == params["id"] {
			tempitem.ID = v.ID
			items[k] = tempitem
			json.NewEncoder(w).Encode(tempitem)
			return
		}
	}

	json.NewEncoder(w).Encode(&List{})

}

func server() {
	r := mux.NewRouter()
	r.HandleFunc("/list", getLists).Methods("GET")
	r.HandleFunc("/list/{id}", getList).Methods("GET")
	r.HandleFunc("/list", createList).Methods("POST")
	r.HandleFunc("/list/{id}", updateList).Methods("PUT")
	r.HandleFunc("/list/{id}", deleteList).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":5000", r))
}

func main() {
	server()
}
