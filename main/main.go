package main

import (
	//helloservice "GolangBasicHTTP/hello_sayer"
	//"fmt"
	ws "GolangBasicHTTP/welcomestring"
	//"html"
	"log"
	"net/http"
)

func main() {
	log.Println("Server Started...")
	http.HandleFunc("/", ws.WelcomeString)
	log.Fatal(http.ListenAndServe(":8081", nil))
	http.HandlerFunc("/")
}
