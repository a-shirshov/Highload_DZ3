package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)


func slowReply(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1000 * time.Millisecond)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server responsed to you my friend!"))
}

func main() {
	r := mux.NewRouter()
	rApi := r.PathPrefix("/api").Subrouter()
	rApi.HandleFunc("",slowReply).Methods("GET")
	fmt.Println("Server started")
	err := http.ListenAndServe(":8080",r)
	if err != nil {
		log.Fatalln("Error: ",err)
	}
}