package main


import (
	"fmt"
	"net/http"
)

func main(){
	fmt.Println("Welcome to the API server!")

	// Initialize the API server
	initializeAPIServer()
}

func initializeAPIServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've reached the API server!")
	})

	http.ListenAndServe(":8080", nil)
}