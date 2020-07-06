package main

/*		--------TO TEST---------
		curl http://localhost:7777/
		curl -v http://localhost:7777/                            It shows what is coming in header.
*/

import (
	"fmt"
	"log"
	"net/http"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "7777"
)

func main() {
	/*
			http.HandleFunc("/", hellosayer): Here, we are registering the helloWorld function with
		the / URL pattern using HandleFunc of the net/http package, which means helloWorld
		gets executed, passing (http.ResponseWriter, *http.Request) as a parameter to it
		whenever we access the HTTP URL with pattern /.

	*/
	http.HandleFunc("/", hellosayer) // Here we're registering
	/*
			err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil): Here, we are calling
		http.ListenAndServe to serve HTTP requests that handle each incoming connection
		in a separate Goroutine. ListenAndServe accepts two parameters—server address and
		handler. Here, we are passing the server address as localhost:8080 and handler as
		nil, which means we are asking the server to use DefaultServeMux as a handler.

	*/

	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatalln("Error while starting server : ", err)
		return
	}
}
func hellosayer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Aliens.")
}
