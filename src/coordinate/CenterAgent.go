package coordinate

import (
    "launchpad.net/gozk"
    "github.com/garyburd/redigo/redis"
    "errors"
    "strconv"
    "time"
    "fmt"
)

var connCount int = 1

const (
    COUNT_SCRIPT string = `
        local circle, begin, min, max, step = tonumber(ARGV[1]), tonumber(ARGV[2]), tonumber(ARGV[3]), tonumber(ARGV[4]), tonumber(ARGV[5]) 
        if 0 == step or 0 > max or 0 > min or 0 > begin then return 0 end 
        if 1 == circle and min > begin then return 0 end 
        if -1 == circle and min > begin then return 0 end 
        local numberNow = tonumber(redis.call('get',KEYS[1]))
        if not numberNow then
            redis.call('set',KEYS[1],begin)
            numberNow = begin
        end
        local nextTmp = numberNow + step 
        if 1 == circle then
            if nextTmp > max then nextTmp = begin+nextTmp%(max-begin) end
            redis.call('set',KEYS[1],nextTmp)
            return nextTmp
        end
        if -1 == circle then
            if nextTmp < min then return 0 end
            redis.call('set',KEYS[1],nextTmp)
            return nextTmp
        end
        return 0;
        `
)

type CenterAgent struct{
    ZkConn *zookeeper.Conn
    RedisPool *redis.Pool
    RedisCountScript *redis.Script
}
/*
func NewCenterAgent(zkAddr string, redisAddr string) *CenterAgent {
    zconn, session, zerr := zookeeper.Dial(zkAddr, 5e9)
    if zerr != nil {
        return nil
    }
    event := <-session
    if event.State != zookeeper.STATE_CONNECTED {
        return nil
    }
    rconn, rerr := redis.DialTimeout("tcp", redisAddr, 0, 1*time.Second, 1*time.Second)
    if rerr != nil {
        return nil
    }
    s := redis.NewScript(1, COUNT_SCRIPT)
    return &CenterAgent{zconn, rconn, s}
} */

func NewCenterPoolAgent(zkAddr string, redisAddr string) *CenterAgent {   
    zconn, session, zerr := zookeeper.Dial(zkAddr, 5e9)
    if zerr != nil {
        return nil
    }
    event := <-session
    if event.State != zookeeper.STATE_CONNECTED {
        return nil
    }
    pool := &redis.Pool{
        MaxIdle:1000,
        MaxActive:5000,
        IdleTimeout:20 * time.Second,
        Dial: func()(redis.Conn,error){
	    connCount = connCount + 1
            fmt.Printf("******** connection " + fmt.Sprintf("%d", connCount) +  " ********")
            return redis.Dial("tcp",redisAddr)
        },
    }
    s := redis.NewScript(1, COUNT_SCRIPT)
    return &CenterAgent{zconn, pool, s}
}

func (z *CenterAgent) GetDataInPath(path string) (data map[string] string, err error) {
    children, _, err := z.ZkConn.Children(path)
    if err != nil {
        return nil, err
    }
    result := make(map[string] string)
    for i:=0; i < len(children); i++ {
        nodeData, _, err := z.ZkConn.Get(path + "/" + children[i])
        if err != nil {
            continue
        }
        result[children[i]] = nodeData
    }
    return result, nil
}

func (z *CenterAgent) CountNum(key string, circle bool, begin int64, min int64, max int64, step int16) (numberNext string, err error) {
    if "" == key || 0 == step  || 0 == begin {
        return "", errors.New("param_error")
    }

    ifCircle := -1
    if circle {
        ifCircle = 1
    }
    c := z.RedisPool.Get()
    v, err := z.RedisCountScript.Do(c, key, ifCircle, begin, min, max, step)
    defer c.Close()
    if err != nil {
       fmt.Println(err)
       fmt.Printf("key:%s ifCircle:%d begin:%d min:%d max:%d step:%d \n",key,ifCircle,begin,min,max,step)
       return  "", err
    }
    switch value := v.(type) {
        case int64:
            sv := strconv.FormatInt(value, 10)
            return sv, nil
        case string:
            return value, nil
        default:
            return "", errors.New("Failed to judge the type of value that lua script returned")
    }

}

func (z *CenterAgent) close() {
   // z.ZkConn.Close()
   // z.RedisPool.Close()
}

