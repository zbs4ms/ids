package coordinate_test

import (
    "coordinate"
    "testing"
    "strconv"
    "time"
)

var centerAgent = coordinate.NewCenterAgent("zk_test:2181", "redis_test:6379")

func TestGetDataInPath(t *testing.T) {
    if centerAgent == nil {
        t.Errorf("Faild to connect zk_test:2181 or redis_test:6379")
    }

    data, err := centerAgent.GetDataInPath("/")
    if err != nil || nil == data {
        t.Errorf("Faild to get data from /")
    }

}

func TestCountNum(t *testing.T) {
    if centerAgent == nil {
        t.Errorf("Faild to connect zk_test:2181 or redis_test:6379")
    }
    
    var numberNow, numberNext int64
    numberNow = 5 
    for i := 0; i < 3; i++ {
        v, err := centerAgent.CountNum(strconv.FormatInt(time.Now().Unix(), 10), false, 5, 0, 99999, 2)
        if err != nil {
            t.Errorf("Failed to count by coordinator: " + err.Error())
        }
        numberNext, _ = strconv.ParseInt(v, 10, 64)
        if numberNow == numberNext {
            continue
        }
        if numberNext - numberNow != 2 {
            t.Errorf("Failed to get wrong number %d and %d", numberNow, numberNext)
        }
        numberNow = numberNext
    }
}
