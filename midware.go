package main

import (
	"fmt"
	"net/http"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Logging: Incoming request to /hello")
		next.ServeHTTP(w, r)
	})
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, this a middleware, from the /hello endpoint!"))
}

func main() {
	router := http.NewServeMux()

	router.Handle("/hello", LoggerMiddleware(http.HandlerFunc(HelloHandler)))

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", router)
}