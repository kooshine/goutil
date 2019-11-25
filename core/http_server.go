package core

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/kooshine/goutil/db"
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
    r.HandleFunc("/home", svr.test_handler)
    server.Handle("/", r)

    go func() {
        fmt.Println("start listen http...")
        err := http.ListenAndServe(svr.Host, server)
        if err != nil {
            log.Fatal("HttpServer Listen Failed:", err)
       }
    }()

    <- exit
}

func (svr *HttpServer) test_handler(res http.ResponseWriter, req *http.Request) {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println(err)
            res.WriteHeader(http.StatusInternalServerError)
        }
    }()
    req.ParseForm()
    uid := req.Form["uid"]
    fmt.Println(uid)

    if req.Method == "GET" {
        for k, v := range req.Form {
            fmt.Println("key:", k)
            fmt.Println("val:", v)
        }
        rows := db.SQLManager.Count("info_website")
        fmt.Println(rows)

        result := db.SQLManager.Search("info_website", "*", "id > -1")
        ret, _ := json.Marshal(result)
        res.Write([]byte(string(ret)))
    }else if req.Method == "POST" {
        buf := make([]byte, 1024)
        n, _ := req.Body.Read(buf)
        fmt.Println("raw data:", string(buf[:n]))
    }else if req.Method == "PUT" {
        fmt.Println(req.Method)
    }else if req.Method == "DELETE" {
        fmt.Println(req.Method)
    }
}
