package count_test


import (
    "testing"
    "coordinate"
    "count"
)

var centerAgent = coordinate.NewCenterAgent("192.168.13.157:50001", "192.168.13.157:60000")


func TestSequentialCount32(t *testing.T){
    if nil == centerAgent {
        t.Errorf("Failed to connect zk_test:2181 or redis_test:6379")
    }
    numberNextMax, err := count.SequentialCount(centerAgent, "php-biz", 1, 0, 9999, 3)
    if err != nil {
        t.Errorf("Failed to count sequential: " + err.Error())
    }

    t.Errorf(numberNextMax)
}

func TestSequentialCount64(t *testing.T) {
    if nil == centerAgent {
        t.Errorf("Failed to connect zk_test:2181 or redis_test:6379")
    }
    numberNextMin, err := count.SequentialCount(centerAgent, "cpp-biz", -1, 9999, 0, 3)
    if err != nil {
        t.Errorf("Failed to count sequential: " + err.Error())
    }

    t.Errorf("64: " +  numberNextMin)
    
}
