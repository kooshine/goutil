package core

import (
    "net"
    "net/http"
)

type HttpServer struct {
    Host string
    Name string
}

func init() {
    HttpConnPool = make(map[string]*HttpConn)
}
