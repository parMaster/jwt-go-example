package main

import (
	"log"
	"net/http"
)

func main() {
	// "Signin" and "Welcome" are the handlers that we will implement
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/welcome", Welcome)
	http.HandleFunc("/refresh", Refresh)
	http.HandleFunc("/logout", Logout)

	// Serving web pages
	http.Handle("/", http.FileServer(http.Dir("./web")))

	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))
}
