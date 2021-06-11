package main

import (
	"net/http"
)

func main(){
	http.HandleFunc("/users", usersHandler)
	go http.ListenAndServe("0.0.0.0:8080", nil)
	go http.ListenAndServe("0.0.0.0:8081", nil)
	for{

	}
}
