package services

import (
    "ttmv/dc/uuid"
    "count"
    "global"
    "strconv"
)

var resultNoauthority = &uuid.UUIDResponse{uuid.ResponseCode_ERROR, "no authority", nil}
var resultErrorUnknown = &uuid.UUIDResponse{uuid.ResponseCode_UNKNOWN, "error inside system", nil}


type UUIDGenerator struct{}

func checkRequest(bizCode string, token string, num int16) *uuid.UUIDResponse {
    if num < 1 || num > 50 {
        return &uuid.UUIDResponse{uuid.ResponseCode_ERROR, "param num invalid", nil}
    }
    if !Verify(bizCode, token) {
        return resultNoauthority
    }
    //if !Permission("0101", bizCode) {
    //     return resultNoauthority
    // }
    return nil
}

func (u *UUIDGenerator) GetCurrentID(bizCode string, token string, num int16) (r *uuid.UUIDResponse, err error){
    check := checkRequest(bizCode, token, num)
    if check != nil {
        return check, nil
    }
    
    result := make([]string, num)
    var i int16
    for i = 0; i < num ; i++ {
        number := count.Random64(global.QueCounter, global.MachineCode)
        if nil == &number || "" == number {
            return resultErrorUnknown, nil
        }
        result[i] = number
    }

    return &uuid.UUIDResponse{uuid.ResponseCode_OK, "", result}, nil
}

func (u *UUIDGenerator) GetPrimaryID32(bizCode string, token string, max int64, num int16) (r *uuid.UUIDResponse, err error){
    check := checkRequest(bizCode, token, num)
    if check != nil {
        return check, nil
    }
    numberNextMax, err := count.SequentialCount(global.CoordinateCenter, bizCode, 1, 0, max, num)
    if err != nil {
        return resultErrorUnknown, nil
    }

    result := make([]string, num)
    v, err := strconv.ParseInt(numberNextMax, 10, 64)
    var i int16
    for i = 0; i < num; i++ {
        result[i] = strconv.FormatInt(v - (int64(num) - int64(i) -1), 10)
    }

    return &uuid.UUIDResponse{uuid.ResponseCode_OK, "", result}, nil
}

func (u *UUIDGenerator) GetPrimaryID64(bizCode string, token string, min int64, num int16) (r *uuid.UUIDResponse, err error){
    check := checkRequest(bizCode, token, num)
    if check != nil {
        return check, nil
    }

    numberNextMin, err := count.SequentialCount(global.CoordinateCenter, bizCode, -1, min, 0, num)
    if err != nil {
        return resultErrorUnknown, nil
    }

    result := make([]string, num)
    v, err := strconv.ParseInt(numberNextMin, 10, 64)
    if err != nil {
        return resultErrorUnknown, nil
    }
    var i int16
    for i = 0; i < num; i++ {
        result[i] = strconv.FormatInt(v + (int64(num) - int64(i) -1), 10)
    }

    return &uuid.UUIDResponse{uuid.ResponseCode_OK, "", result}, nil
}

func (u *UUIDGenerator) GetOrder(bizCode string, token string, limit int64, num int16) (r *uuid.UUIDResponse, err error){
    check := checkRequest(bizCode, token, num)
    if check != nil {
        return check, nil
    }
    numberNextMax, err := count.CircleCount(global.CoordinateCenter, bizCode,  limit, num)
    if err != nil {
        return resultErrorUnknown, nil
    }
    result := make([]string, num)
    v, err:= strconv.ParseInt(numberNextMax, 10, 64)
    if err != nil {
        return resultErrorUnknown, nil
    }
    var i int16
    var n int64 = int64(v)
    for i = 0; i < num; i++ {
        n = n - 1
        if n <= 0 {
            n = limit
        }
        //result[i] = strconv.FormatInt(v - (int64(num) - int64(i) -1), 10)
        result[i] = strconv.FormatInt(n,10)
    }
    return &uuid.UUIDResponse{uuid.ResponseCode_OK, "", result}, nil
}
