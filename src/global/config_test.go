package global

import(
   "fmt"
   "testing"
   "global"
)

func TestConfig(t *testing.T){
    str:err := Get("uuidCluster","machineCode")
    if err !=nil{
      fmt.Println("error")
    }
    fmt.Println(str)
}
