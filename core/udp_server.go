package core

import (
    "fmt"
    "net"
)

type UdpServer struct {
    Host string
    Name string
}

func (svr *UdpServer) Run() {
    exit := make(chan int, 1)

    if svr.Host == "" {
        return
    }

    var addr *net.UDPAddr
    addr, _ = net.ResolveUDPAddr("udp4", svr.Host)

    udpListener, err := net.ListenUDP("udp4", addr)
    if err != nil {
        fmt.Println("udp error")
    }
    defer udpListener.Close()

    go func() {
        fmt.Println("udp conn>>>")
        for {
            data := make([]byte, 4098)
            n, remoteAddr,err := udpListener.ReadFromUDP(data)
            if err != nil {
                fmt.Println("read udp err")
                continue
            }
            go svr.ReceiveData(udpListener, remoteAddr, data[:n])
        }
    }()
    <-exit
}

func (svr *UdpServer) ReceiveData(udpListener *net.UDPConn, remoteAddr *net.UDPAddr, data []byte) {
    fmt.Printf("Accpet UDP Connection, Receive Data, remoteAddr: %s, len=%d, data(hex) = [ %x ]\n", remoteAddr, len(data), data)
    c := newConn("udp", remoteAddr.String(), svr.Name)
    c.inBufs = data
    go c.Start()
}
