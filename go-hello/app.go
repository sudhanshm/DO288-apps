package main

import (
    "net/http"
    "flag"
    "fmt"
)

var lang = flag.String("lang", "en", "run app with language support - default is english")

func main() {
    var port = "8080"
    http.HandleFunc("/", rootHandler)
    fmt.Printf("Starting server on port %v...\n", port)
    http.ListenAndServe(":"+port, nil)
}

func rootHandler(response http.ResponseWriter, request *http.Request) {

    flag.Parse()

    fmt.Fprintf(response, "Hola user1!. Bienvenido!\n", request.URL.Path[1:])
}
