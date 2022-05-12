package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)


func slowReply(w http.ResponseWriter, r *http.Request) {
	i := time.Duration(rand.Intn(1000))
	time.Sleep(i * time.Millisecond)
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