package yudian

import (
    //"github.com/kooshine/goutil/device/protocol"
)

type Yudian_Protocol struct {
    //protocol.BaseProtocol
    AppType string
    DevId int64
}


func (yd *Yudian_Protocol) ParseDevType() (string) {
    return "yudian_elec"
}

func (yd *Yudian_Protocol) ParseDevId(data []byte) (output int64) {
    devId := data[8:14]
    for i := 0; i < len(devId); i++ {
        output = (output <<8) + int64(uint32(devId[i]))
    }
    return output
}
