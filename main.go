package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

func main(){

	router := mux.NewRouter()
	router.HandleFunc("/",info).Methods("GET")
	log.Fatal(http.ListenAndServe(":9999",router))
}

func info(w http.ResponseWriter, req *http.Request){
	w.Write([]byte("Amazing web page!!"))
}