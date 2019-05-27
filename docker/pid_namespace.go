/***********************************************************************
# File Name: pid_namespace.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2018-12-04 15:41:34
*********************************************************************/
package main

import(
    "log"
    "os"
    "os/exec"
    "syscall"
)

func main(){
    cmd := exec.Command("sh")
    cmd.SysProcAttr = &syscall.SysProcAttr{
        Cloneflags:syscall.CLONE_NEWUTS|syscall.CLONE_NEWIPC|syscall.CLONE_NEWPID,
    }

    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run();err != nil{
        log.Fatal(err)
    }
}
