/***********************************************************************
# File Name: storage.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2018-12-25 14:05:26
*********************************************************************/
package storage

import(
    "fmt"
    "log"
    "net/smtp"
)

func bytesInUse(username string) int64{ return 0}

const sender = "notification@example.com"
const password = "correcthorsebatterystaple"
const hostname = "smtp.example.com"

const template = `Warning: you are using %d bytes of storage, %d%% of your quota.`o

func CheckQuota(username string) {
    used := bytesInUser(username)
    const quota = 1000000000
    percent := 100 * userd / quota
    if percent < 90{
        return 
    }
    msg := fmt.Sprintf(template,used,percent)
    notifyUser(username,msg)
}

var notifyUser = func(username,msg string){
    auth := smtp.PlainAuth("",sender,password,hostname)
    err := smtp.SendMail(hostname + ":587",auth,sender,[]string{username},[]byte(msg))
    if err != nil{
        log.Printf("smtp.SendMail(%s) failed:%s",username,err)
    }
}
