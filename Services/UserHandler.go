package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)


func usersHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case http.MethodGet: getHandler(w, r)
	case http.MethodPost: postHandler(w, r)
	}
}

func getHandler(w http.ResponseWriter, r *http.Request){
	res, _ := json.Marshal(Users)
	w.Header().Add("Content-Type", "application/json")
	w.Write(res)
}


func postHandler(w http.ResponseWriter, r *http.Request){
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
