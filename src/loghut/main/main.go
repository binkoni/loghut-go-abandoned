package main

import "fmt"
import "net/http"
import "loghut/controller"
import "syscall"

func main() {
    fmt.Println(syscall.Getpid())
    http.HandleFunc("/", controller.Handle)
    http.ListenAndServe("127.0.0.1:8080", nil)
}
