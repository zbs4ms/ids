package services_test

import (
    a  "services"
    "testing"
    "global"
    c "coordinate"
)

func TestAuthorityWithZk(t *testing.T) {
    coor := c.NewCenterAgent("192.168.13.157:50001", "192.168.13.157:60000")

    if nil == coor {
        t.Errorf("Failed to init coordinator with ttmv_uuid_zk:2181 and ttmv_uuid_redis:6379")
    }

    err := a.InitAuthority(coor)
    if err != nil {
        t.Errorf("Failed to get data from %s or %s", global.VERIFY_INFO_PATH, global.PERMISSION_INFO_PATH)
    }

    if !a.Verify("CID64", "123456") {
        t.Errorf("Failed to verify \"bizCode\" ")
        for key, value := range global.LegalUser {
            t.Errorf(key + ":" + value)
        }
    }

    if !a.Permission("0101", "CID64") {
        t.Errorf("Failed to give \"bizCode\" permission \"0101\" ")
        for key, value := range global.Permission {
            t.Errorf(key + ":" + value)
        }
    }
}
