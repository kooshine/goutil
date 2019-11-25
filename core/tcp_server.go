package core

import (
    "fmt"
    "net"
)

type TcpServer struct {
    Host string
    Name string
}

func (svr *TcpServer) Run() {
    exit := make(chan int,1)

    if svr.Host == "" {
        return
    }

    var addr *net.TCPAddr
    addr, _ = net.ResolveTCPAddr("tcp", svr.Host)

    listener, err := net.ListenTCP("tcp", addr)
    if err != nil {
        fmt.Println("error")
    }
    defer listener.Close()

    go func() {
        for {
            conn, err := listener.AcceptTCP()
            if err != nil {
                fmt.Println("error")
            }
            fmt.Println(svr.Name, conn.RemoteAddr().String())
        }

        exit <- 1
    }()

    fmt.Println("Tcp Server start on >>>")
    <- exit
}
