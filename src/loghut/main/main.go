package main

import "net/http"
import "loghut/controller"

func main() {
    http.HandleFunc("/", controller.Handle)
    http.ListenAndServe("127.0.0.1:8080", nil)
}
