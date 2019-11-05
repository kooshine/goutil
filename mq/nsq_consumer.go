package nsq
import (
    "fmt"
    "encoding/json"
    _ "time"
    "errors"
    "github.com/nsqio/go-nsq"
    "github.com/kooshine/goutil/db"
)

type Nsq_consumer struct {
    c *nsq.Consumer
    tag string
}
type Recv struct {
    Msg string
}

func (r *Nsq_consumer) HandleMessage(msg *nsq.Message) error {
    fmt.Println("receive: ", msg.Body)

    d := &Recv{}
    json.Unmarshal(msg.Body, &d)
        /*
    if err := json.Unmarshal(msg.Body, &d); err == nil {
        fmt.Println(">>>>>>>>>>>>>>>>>>",d.Msg)
        return err
    }*/
    fmt.Println(">>>>>>>>>>>>>>>>>>",d.Msg)
    var dataMap = map[string]string {
        "topic": "test",
        "channel": "channel",
        "message": d.Msg,
    }
    ret := db.SQLManager.Insert("packet", dataMap)
    fmt.Println(ret)
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
    nsq_c := &Nsq_consumer {
        c:  c,
        tag: "tag",
    }
    c.AddHandler(nsq_c)
    fmt.Println(c)

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
