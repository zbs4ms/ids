package count_test

import (
    "testing"
    "count"
)


func TestRandom64(t *testing.T) {

    queCounter := count.NewQueCounter(1, 100000000, 2)

    chs := make([]chan [1000]string, 100)
    for i := 0; i < 100; i++ {
        chs[i] = make(chan [1000]string)
        go calcRandom(chs[i], queCounter)
    }

    var all map[string] string = make(map[string] string)
    for _, ch := range(chs) {
        orders := <- ch
        for i:=0; i < 1000; i++ {
            value, ok := all[orders[i]]
            if ok {
                t.Errorf("Failed for create same id: %s", value)
            }else{
                all[orders[i]] = "1"
            }
        }
    }
}

func calcRandom(ch chan [1000]string, queCounter *count.QueCounter) {
    var orders [1000]string
    for i := 0; i < 1000; i++ {
        orders[i] = count.Random64(queCounter,"CID64")
    }
    ch <- orders
}
