package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"time"
)

func getUsers(){
	for {
		time.Sleep(time.Second)
		res, err := http.Get("http://localhost:8080/users")
		if err != nil {
			log.Fatal(err)
			return
		}
		scanner := bufio.NewScanner(res.Body)
		for i := 0; scanner.Scan() && i < 5; i++ {
			fmt.Println(scanner.Text())
		}
		res.Body.Close()
	}
}


func main() {
	getUsers()
	//for {}
}
