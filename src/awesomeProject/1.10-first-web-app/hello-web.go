package main

import (
	"log"
	"net/http"
)

// 1.- Listen to the root path using the net/http package
// 2.- Write out the hardcoded HTML when request is made
// 3.- Start a web server on port 8080 using the ListenAndServe method

func main() {
	// maps the path pattern / to the function we pass as the second argument
	// when the user hits localhost:8080 the function will be executed
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		// common way of handling HTTP requests throughout the Go standard library
		_, _ = w.Write([]byte (`
			<html>
				<head>
					<title>Chat</title>
				</head>
				<body>
					Let's chat!
				</body>
			</html>
		`))
	})
	//start the web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
