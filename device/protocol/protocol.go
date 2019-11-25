package protocol


type Protocoler interface {
    ParseDevType()(devType string)
    ParseDevId(data []byte)(devId int64)
}

type BaseProtocol struct {
    appType string
    devId int64
}
