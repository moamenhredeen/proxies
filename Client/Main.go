package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"sync"
)

func getUsers(wg *sync.WaitGroup){
	defer wg.Done()
	res, err := http.Get("http://localhost/users")
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


func main() {
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0 ; i < 100 ; i++{
		go getUsers(&wg)
	}
	wg.Wait()
	//for {}
}
