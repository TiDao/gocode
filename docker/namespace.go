/***********************************************************************
# File Name: namespace.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2018-12-04 17:45:37
*********************************************************************/
package main

import(
    "log"
    "os"
    "os/exec"
    "syscall"
)

func main() {
    cmd := exec.Command("bash")
    cmd.SysProcAttr = &syscall.SysProcAttr{
        Cloneflags: syscall.CLONE_NEWUTS|
        syscall.CLONE_NEWIPC|
        syscall.CLONE_NEWPID|
        syscall.CLONE_NEWNS|
        syscall.CLONE_NEWUSER|
        syscall.CLONE_NEWNET,
        UidMappings:[]syscall.SysProcIDMap{{ContainerID:1234,HostID:0,Size:1}},
        GidMappings:[]syscall.SysProcIDMap{{ContainerID:1234,HostID:0,Size:1}},
    }

    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err:=cmd.Run(); err!=nil{
        log.Fatal(err)
    }
    os.Exit(-1)
}
