package controller

import "fmt"
import "net/http"
import "sync"

var actionHandlers map[string]func(*http.ResponseWriter, *http.Request) = make(map[string]func(*http.ResponseWriter, *http.Request))
var mutex *sync.Mutex

func init() {
    mutex = new(sync.Mutex)
    actionHandlers["backup"] = func(responseWriter *http.ResponseWriter, request *http.Request) {
        responseHeader := (*responseWriter).Header()
        responseHeader.Set("Content-Type", "text/html")
        fmt.Fprintln(*responseWriter, "backup<br/>")
    }
    actionHandlers["creation_form"] = func(responseWriter *http.ResponseWriter, request *http.Request) {
        responseHeader := (*responseWriter).Header()
        responseHeader.Set("Content-Type", "text/html")
        fmt.Fprintln(*responseWriter, "creation_form<br/>")
    }
    actionHandlers["create"] = func(responseWriter *http.ResponseWriter, request *http.Request) {
        responseHeader := (*responseWriter).Header()
        responseHeader.Set("Content-Type", "text/html")
        fmt.Fprintln(*responseWriter, "create<br/>")
    }
    actionHandlers["default"] = func(responseWriter *http.ResponseWriter, request *http.Request) {
        responseHeader := (*responseWriter).Header()
        responseHeader.Set("Content-Type", "text/html")
        fmt.Fprintln(*responseWriter, "default<br/>")
    }
    actionHandlers["delete"] = func(responseWriter *http.ResponseWriter, request *http.Request) {
        responseHeader := (*responseWriter).Header()
        responseHeader.Set("Content-Type", "text/html")
        fmt.Fprintln(*responseWriter, "delete<br/>")
    }
    actionHandlers["login"] = func(responseWriter *http.ResponseWriter, request *http.Request) {
        responseHeader := (*responseWriter).Header()
        responseHeader.Set("Content-Type", "text/html")
        fmt.Fprintln(*responseWriter, "login<br/>")
    }
    actionHandlers["logout"] = func(responseWriter *http.ResponseWriter, request *http.Request) {
        responseHeader := (*responseWriter).Header()
        responseHeader.Set("Content-Type", "text/html")
        fmt.Fprintln(*responseWriter, "logout<br/>")
    }
    actionHandlers["modification_form"] = func(responseWriter *http.ResponseWriter, request *http.Request) {
        responseHeader := (*responseWriter).Header()
        responseHeader.Set("Content-Type", "text/html")
        fmt.Fprintln(*responseWriter, "modification_form<br/>")
    }
    actionHandlers["modify"] = func(responseWriter *http.ResponseWriter, request *http.Request) {
        responseHeader := (*responseWriter).Header()
        responseHeader.Set("Content-Type", "text/html")
        fmt.Fprintln(*responseWriter, "modify<br/>")
    }
    actionHandlers["search"] = func(responseWriter *http.ResponseWriter, request *http.Request) {
        responseHeader := (*responseWriter).Header()
        responseHeader.Set("Content-Type", "text/html")
        fmt.Fprintln(*responseWriter, "search<br/>")
    }
}

func Handle(responseWriter http.ResponseWriter, request *http.Request) {
    mutex.Lock()
    if action, ok := request.URL.Query()["action"]; ok {
        if _, ok := actionHandlers[action[0]]; ok {
            actionHandlers[action[0]](&responseWriter, request)
        } else {
            actionHandlers["default"](&responseWriter, request)
        }
    } else {
        actionHandlers["default"](&responseWriter, request)
    }
    mutex.Unlock()
}

