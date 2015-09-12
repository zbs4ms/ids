package global

import (
    "coordinate"
    "count"
)

/*
legalUser 用于接入方身份验证：key为接入方被分配的bizCode, value为接入方被授予的token
authority 用于确定接入方的权限: key为接口业务编号，value为接入方被分配的bizCode
filter 用于生成id过滤器：key为接入方被分配的bizCode， value为过滤器数据
machineCode 为本uuid实例的被分配的机器码
*/
var (
    LegalUser map[string] string = make(map[string] string)
    Permission map[string] string = make(map[string] string)

    IDFilter map[string] []byte = make(map[string] []byte)

    MachineCode string
    ServeAddr string

    CoordinateCenter *coordinate.CenterAgent
    QueCounter *count.QueCounter
)
