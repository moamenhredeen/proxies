package main

import (
	"net/http"
)

func main(){
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/todolist", todolistHandler)
	http.ListenAndServe("0.0.0.0:8080", nil)
}
