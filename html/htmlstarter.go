package main

import (
	"encoding/json"
	"fmt"
	"nats-service/db"
	"nats-service/memory"
	"net/http"
	"strconv"
)

var cache *memory.Cache

func main() {
	CreateCache()
	http.HandleFunc("/getData", getDataHandler)
	http.Handle("/", http.FileServer(http.Dir("static")))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("%s\n", "Can't listen to port")
	}
}

func CreateCache() {
	db.Connect()
	response := db.GetDBData()
	cache = memory.NewCache(response)
}

func getDataHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data := cache.Read(id)
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
