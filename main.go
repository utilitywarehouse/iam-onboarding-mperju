package main

import (
	"fmt"
	"net/http"
)

func Test(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Ping")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/text")
	w.Write([]byte("Success"))

}

func main() {

	fmt.Println("Launching the server...")
	http.HandleFunc("/", Test)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error running the server")
	}
}
