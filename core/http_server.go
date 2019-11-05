package core

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

type HttpServer struct {
    Host string
    Name string
}


func (svr *HttpServer) Run() {
    exit := make(chan int,1)

    server := http.NewServeMux()
    if svr.Host == "" {
        return
    }

    r := mux.NewRouter()
    r.HandleFunc("/test", svr.test_handler)
    server.Handle("/", r)

    go func() {
        fmt.Println("test_http")
        err := http.ListenAndServe(svr.Host, server)
        if err != nil {
            log.Fatal("HttpServer Listen Failed:", err)
       }
    }()

    <- exit
}

func (svr *HttpServer) test_handler(res http.ResponseWriter, req *http.Request) {
    fmt.Println("abc")
}
