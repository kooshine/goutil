package core

import (
    "fmt"
    "time"
    "github.com/kooshine/goutil/device/protocol"
    "github.com/kooshine/goutil/device/yudian"
)

const (
    timeout = 60*60
    defaultBufferSize   = 1024
)

type Conn struct {
    appType string
    connType string
    remoteAddr string
    inBufs []byte
    devId int64
    //parser interface {}
    proto protocol.Protocoler
}


func newConn(connType string, remoteAddr string, name string) *Conn {
    c := &Conn {
        appType:    name,
        connType:   connType,
        remoteAddr: remoteAddr,
        inBufs:     make([]byte, defaultBufferSize),
        devId:      0,
    }
    return c
}

func (c *Conn) Start() {
    defer func() {
        fmt.Println("exit...")
    }()

    go c.heartBeating(timeout)
    fmt.Printf("data= [ % x ]\n", c.inBufs)
    fmt.Printf("appType: %s, connType: %s, remoteAddr:%s\n", c.appType, c.connType, c.remoteAddr)
    if c.appType == "yudian" {
        c.proto = &yudian.Yudian_Protocol {
            AppType:    "yudian",
            DevId:    0,
        }
        devType := c.proto.ParseDevType()
        fmt.Println(devType)
    }
}

func (c *Conn) heartBeating(timeout int){
    time_ticket := time.NewTimer(time.Second *time.Duration(timeout))
    for {
        select {
        case <- time_ticket.C:
            fmt.Println("exit...timeout")
        }
    }
}
