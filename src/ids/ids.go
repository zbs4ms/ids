package main

import (
    "oso/inf/ids"
    "fmt"
    "os"
    "git.apache.org/thrift.git/lib/go/thrift"
    "services"
    "global"
    "coordinate"
    "count"
)


func main(){
    InitProgram()
    startThriftService()
}

func InitProgram(){
    global.CoordinateCenter = coordinate.NewCenterPoolAgent(global.Get("center", "zookeeperAddr"), global.Get("center", "redisAddr"))
    if nil == global.CoordinateCenter {
        fmt.Println("Can not access to coordintor")
        os.Exit(1)
    }
    global.ServeAddr = global.Get("uuidCluster", "serveAddr")

    machineCode := global.Get("uuidCluster", "machineCode") 
    global.MachineCode = machineCode

    global.QueCounter = count.NewQueCounter(1, 8191, 1)

    services.InitAuthority(global.CoordinateCenter)
}

func startThriftService(){
    transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
    protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

    serverTransport, err := thrift.NewTServerSocket(global.ServeAddr)
    if err != nil {
        fmt.Println("Error!", err)
        os.Exit(1)
    }
    
    handler := &services.UUIDGenerator{}
    processor := uuid.NewGeneratorProcessor(handler)

    server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
    fmt.Println("thrift server in", global.ServeAddr)
    server.Serve()
}
