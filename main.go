package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Post("https://google.com", "", nil)
	log.Println(resp)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "halo!")
	})

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
