package global

import (
    "flag"
    "github.com/larspensjo/config"
    "os"
)
 
var (
    con = os.Getenv("UUIDCONFIG")+"/config.ini"
    configFile = flag.String("configfile",con, "General configuration file")
)

func Get(section string, key string) string {
    flag.Parse()
    cfg, err := config.ReadDefault(*configFile)
    if err != nil {
        return ""
    }
    v, err := cfg.String(section, key)
    if err != nil {
        return ""
    }
    return v
}

func Set(section, key, value string) {
    flag.Parse()
    cfg,_ := config.ReadDefault(*configFile)
    if cfg.AddOption(section,key,value){
       cfg.WriteFile(con,0644,"General configuration file")
    }
}

 
 
 
