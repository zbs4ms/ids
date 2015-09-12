package count

import (
    c "coordinate"
    "errors"
)


func SequentialCount(coor *c.CenterAgent, bizCode string, direction int8, min int64, max int64, num int16) (number string, err error) {
    if nil == coor || "" == bizCode || num < 1{
        return "", errors.New("param_invalid: coor or bizCode or num")
    }
    
    var numberNextMax string
    switch direction {
    case 1:
        numberNextMax, err = coor.CountNum(bizCode, false, 1, 0, max, num)
        return numberNextMax, err
    case -1:
        numberNextMax, err = coor.CountNum(bizCode, false, 18014398509481984, min, 0, 0 - num)
        return numberNextMax, err
    default:
        return "", errors.New("param_invalid: direction")
    }
}



func CircleCount(coor *c.CenterAgent, bizCode string, max int64, num int16) (number string, err error) {
    if nil == coor || "" == bizCode || max < 1 || num < 1{
        return "", errors.New("param_invalid")
    }

    numberNextMax, err := coor.CountNum(bizCode, true, 1,  0, max, num)
    return numberNextMax, err
    
}
