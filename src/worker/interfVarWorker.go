package worker

import (
    "global"
    "coordinate"
)

func StartInterfVarWorker(coor Coordinator) (err error){
    err := initInterfVarWorker(coor)
    return err
}

func initInterfVarWorker(coor Coordinator) error{
    global.verify, err = coor.GetDataInPath(global.VERIFY_INFO_PATH)
    if err != nil {
        return nil
    }
    global.permission, err = coor.GetDataInPath(global.PERMISSION_INFO_PATH)
    if err != nil {
        return nil
    }
}

