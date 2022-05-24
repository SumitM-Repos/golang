package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8081", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {

	context := r.Context()

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "Hello")
	case <-context.Done():
		err := context.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}

}
