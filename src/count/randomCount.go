package count

import (
    "fmt"
    "time"
    "strconv"
)

func Random64(queCounter *QueCounter, machineCode string) string{
    var uuid string
    now := time.Now()
    nano := now.UnixNano()
    codeInt,_ := strconv.Atoi(machineCode)
    timeStamp := nano >>20
    lenInt,_ := strconv.Atoi(fmtToLength(queCounter.NextOrder(),4))
    uuid =fmt.Sprintf("%d",((^(timeStamp<<23)+1)^(int64(codeInt)<<13)^(int64(lenInt))))
    return uuid
}

//todo: 本函数可以优化
func fmtToLength(number uint64, length int) string{
    numberStr := fmt.Sprintf("%d", number)

    gap := length - len(numberStr)
    if gap > 0 {
        for i := 0; i < gap; i++ {
            numberStr = "0" + numberStr
        }
    }

    return numberStr
}

