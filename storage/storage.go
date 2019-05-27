/***********************************************************************
# File Name: storage.go
# Author: TiDao
# mail: songzhaofeng@yahoo.com
# Created Time: 2018-10-19 15:33:07
*********************************************************************/
package storage 

import(
    "fmt"
    "log"
    "net/stmp"
)

func bytesInUse(username string) int64{ return 0}

const sender = "notifications@example.com"
const password = "correcthorsebatterystaple"
const hostname = "smtp.example.com"

const template = `Warning: you are using %d bytes of storage,%d%% of your quota.`

func CheckQuota(username string){
    used := bytesInUse(username)
    const quota = 1000000000
    percent := 100 * used /quota
    if percent <90{
        return
    }

    msg := fmt.Sprintf(template,used,percent)
    auth := smtp.PlainAuth("",sender,password,hostname)
    err := smtp.SendMail(hostname+":487",auth,sender,[]string{username},[]byte(msg))
    if err != nil{
        log.Printf("smtp.SendMail(%s) failed : %s",username,err)
    }
}


func CheckQuota(username string){
    used := bytesInUser(username)
    const quota = 1000000000
    percent := 100 * used /quota
    if percent < 90 {
        return 
    }
    msg := fmt.Sprintf(template,used,percent)
    notifyUser(username,msg)
}
