package nsq
import (
    "fmt"
    _ "time"
    "errors"
    "github.com/nsqio/go-nsq"
)

type nsq_consumer struct {}

func (*nsq_consumer) HandleMessage(msg *nsq.Message) error {
    fmt.Println("receive: ", msg.NSQDAddress, "message:", string(msg.Body))
    return nil
}

func Consumer(topic string, channel string, host string) error {
    fmt.Println("start Conusumer>>>")

    c, err := nsq.NewConsumer(topic, channel, nsq.NewConfig())
    if err != nil {
        err_msg := "Error, InitConusmer failed"
        fmt.Println(err_msg)
        return errors.New(err_msg)
    }
    c.SetLogger(nil, 0)
    c.AddHandler(&nsq_consumer{})

    //建立NSQLookupd连接
    if err := c.ConnectToNSQLookupd(host); err != nil {
        fmt.Println("Lookupd conn failed")
        panic(err)
    }
    /*
    //建立多个nsqd连接
    if err := c.ConnectToNSQDs([]string{"127.0.0.1:4150", "127.0.0.1:4152"}); err != nil {
        panic(err)
    }

    // 建立一个nsqd连接
    if err := c.ConnectToNSQD("127.0.0.1:4150"); err != nil {
        panic(err)
    }
    */
    return nil
}
