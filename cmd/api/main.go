package main


import (
	"fmt"
	"net/http"
  "github.com/gorilla/mux"
)

func main(){
	fmt.Println("Welcome to the API server!")

	// Initialize the API server
	initializeAPIServer()
}

func initializeAPIServer() {
	    router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've reached the API server!")
	})

	http.ListenAndServe(":8080", router)
}