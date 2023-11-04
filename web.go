// Building a Web API with Go

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"encoding/json"
)

// Serving JSON data

type Person struct {
	Id int
	Name string
}

func ReturnJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	p := Person {
		Id: 1,
		Name: "a person",
	}

	json.NewEncoder(w).Encode(p)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello, this is a running web server.\n")
	// w = a writer
}

// Responding to a request

func GetImage(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("fav.webm")

	// Read the entire JPG file into memory
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	// Set the Content Type header
	w.Header().Set("Content-Type", "fav.webm")

	// Write image to the response
	w.Write(content)
}


func main() {
	// do something with http
	http.HandleFunc("/hello", hello,)
	http.HandleFunc("/image", GetImage)
	http.HandleFunc("/json", ReturnJson)
	
	// ServeMux doc, how patterns are matched

	// Starting the server
	http.ListenAndServe(":8090", nil)
}
