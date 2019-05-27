/***********************************************************************
# File Name: uts_namespace.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2018-12-04 15:01:51
*********************************************************************/
package main

import(
    "os/exec"
    "syscall"
    "os"
    "log"
)

func main(){
    cmd := exec.Command("sh")
    cmd.SysProcAttr = &syscall.SysProcAttr{
        Cloneflags: syscall.CLONE_NEWUTS,
    }
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil{
        log.Fatal(err)
    }
}

