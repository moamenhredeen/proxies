package main


import (
"encoding/json"
"io/ioutil"
"log"
"net/http"
)


type TodoItem struct {
	Id int `json:"id"`
	Header string `json:"header"`
	Body string `json:"body"`
}


var TodoItems = []TodoItem{
	TodoItem{1, "netzwerke", "proxy programmieren"},
	TodoItem{2, "projektmanagement", "projekt plan erstellen"},
	TodoItem{3, "simulation technik", "plant simulation ubung machen"},
	TodoItem{4, "mobile systeme", "referat über software architektur"}}

func todolistHandler(w http.ResponseWriter, r *http.Request){
	log.Println(r)
	switch r.Method {
	case http.MethodGet:  todolistGetHandler(w, r)
	case http.MethodPost: todolistPostHandler(w, r)
	}
}

func todolistGetHandler(w http.ResponseWriter, r *http.Request){
	res, _ := json.Marshal(TodoItems)
	w.Header().Add("Content-Type", "application/json")
	w.Write(res)
}


func todolistPostHandler(w http.ResponseWriter, r *http.Request){
	var newTodoItem TodoItem
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(bodyBytes, &newTodoItem)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	TodoItems = append(TodoItems, newTodoItem)
	w.WriteHeader(http.StatusCreated)
	return
}
