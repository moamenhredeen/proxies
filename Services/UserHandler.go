package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)


type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}

type Data struct {
	Users []User
	Service string
}

var Users = []User{
	User{1, "moamen", 20},
	User{2, "rajaa", 23},
	User{3, "saja", 30},
	User{4, "imam", 25}}


func usersHandler(w http.ResponseWriter, r *http.Request){
	log.Println(r)
	switch r.Method {
	case http.MethodGet: userGetHandler(w, r)
	case http.MethodPost: userPostHandler(w, r)
	}
}

func userGetHandler(w http.ResponseWriter, r *http.Request){
	data := Data{Users:Users, Service: os.Getenv("NAME")}
	res, _ := json.Marshal(data)
	w.Header().Add("Content-Type", "application/json")
	w.Write(res)
}


func userPostHandler(w http.ResponseWriter, r *http.Request){
	var newUser User
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(bodyBytes, &newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	Users = append(Users, newUser)
	w.WriteHeader(http.StatusCreated)
	return
}
