package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var id int
	fmt.Print("Enter ID: ")
	fmt.Scan(&id)

	url := fmt.Sprintf("http://localhost:8080/hello-user?id=%d", id)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("failed to get: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read response: %v", err)
	}
	fmt.Println(string(body))
}
