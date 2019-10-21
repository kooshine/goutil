package nsq

import (
    "fmt"
    "strings"
    "github.com/nsqio/go-nsq"
)

type NsqWriter struct {
    Host string
    Producer *nsq.Producer
    Messages chan[2]string
}

var nsqdHosts []string
var nsqers []*NsqWriter

func InitNsq(nsqd string) {
    fmt.Println("Init NSQ >>>>>>")
    nsqdHosts = strings.Split(nsqd, ",")
    nsqers = make([]*NsqWriter, len(nsqdHosts))

    for i, host := range nsqdHosts {
        fmt.Printf("init %d\n", i)
        nsqers[i] = &NsqWriter {
            Host: host,
            Messages: make(chan [2]string),
        }

        go func(idx int) {
            for {
                fmt.Printf("init %d\n", idx)
            }
        }(i)
    }
}

