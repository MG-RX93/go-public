package main

// import all required packages
import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static")) // returns a http handler that serves static files from the ./static dir
	http.Handle("/", fileServer) // handle all requests
	http.HandleFunc("/form", formHandler) // handle requests to /form
	http.HandleFunc("/hello", helloHandler) // handle requests to /hello
}