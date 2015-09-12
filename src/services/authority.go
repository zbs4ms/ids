package services

import (
    "global"
    "strings"
    c "coordinate"
)

func InitAuthority(coor *c.CenterAgent) error{
    verifyInfo, err := coor.GetDataInPath(global.VERIFY_INFO_PATH)
    if err != nil {
        return err
    }
    for key, value := range verifyInfo {
        global.LegalUser[key] = value
    }
    
    
    permissionInfo, err := coor.GetDataInPath(global.PERMISSION_INFO_PATH)
    if err != nil {
        return err
    }
    for key, value := range permissionInfo {
        global.Permission[key] = value
    }

    return nil
}

func Verify(bizCode string, token string) bool {
    if nil == global.LegalUser || nil == &bizCode {
        return false
    }
    
    value, ok := global.LegalUser[bizCode]
    if ok {
        if value == token {
            return true
        }
    }
    return false
}

func Permission(permCode string, bizCode string) bool {
    if nil == global.Permission || nil == &bizCode || "" == bizCode {
        return false
    }

    value, ok := global.Permission[permCode]
    if ok {
        if strings.Index(value, bizCode) != -1 {
            return true
        }
    }
    return false
}
