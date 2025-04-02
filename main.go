package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

// type route struct {
// 	path, method string
// 	handleFunc http.Handler
// }



func main() {
	// ServeMux will tie the base url 'localhost
	mux := http.NewServeMux()

	// var routes = []route{
	// 	newRoute("GET", "/", )
	// }
	// creating routes
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)
	mux.HandleFunc("/data", postData)

	// Either errors or spins up server on PORT
	err := http.ListenAndServe(":8000", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed..\n")
	} else if err != nil {
		fmt.Printf("error because: %v", err)
		os.Exit(1)
	}
}

// func newRoute(method string, path string, handler http.HandlerFunc) route {
// 	return route{method, path, handler}
// }

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("goot / req\n")
	io.WriteString(w, "ROOOOOOT\n")

}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func postData(w http.ResponseWriter, r * http.Request) {
	fmt.Printf("%v / request\n", r.Method)
	if r.Method == "POST" {
		fmt.Println(r.Body)
		body, err := io.ReadAll(r.Body)
		derr := json.NewDecoder(r.Body)
		fmt.Println("deerrrr", derr)
		if err != nil {
			os.Exit(3)
		}
		fmt.Println(string(body))
		
	}
}