/***********************************************************************
# File Name: cgroup.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2018-12-05 17:10:46
*********************************************************************/
package main

import(
    "os/exec"
    "path"
    "os"
    "fmt"
    "io/ioutil"
    "syscall"
    "strconv"
    "log"
)

const cgroupMemoryHierarchyMount = "/sys/fs/cgroup/memory"

func main(){
    if os.Args[0] == "/proc/self/exe" {
        fmt.Printf("current pid %d ",syscall.Getpid())
        fmt.Println()
        cmd := exec.Command("sh","-c","stress --vm-bytes 200m --vm-keep 1")
        cmd.SysProcAttr = &syscall.SysProcAttr{
        }
        cmd.Stdin = os.Stdin
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        if err := cmd.Run(); err !=nil{
            log.Fatal(err)
            os.Exit(1)
        }
    }

    cmd := exec.Command("/proc/self/exe")
    cmd.SysProcAttr = &syscall.SysProcAttr {
        Cloneflags:syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
    }
    if err:=cmd.Start(); err!=nil{
        log.Fatal(err)
        os.Exit(1)
    } else {
        fmt.Printf("%v\n",cmd.Process.Pid)
        os.Mkdir(path.Join(cgroupMemoryHierarchyMount,"testmemorylimit"),0755)
        ioutil.WriteFile(path.Join(cgroupMemoryHierarchyMount,"testmemorylimit","tasks"),[]byte(strconv.Itoa(cmd.Process.Pid)),0644)
        ioutil.WriteFile(path.Join(cgroupMemoryHierarchyMount,"testmemorylimit","memory.limit_in_bytes"),[]byte("100m"),0644)
    }
    cmd.Process.Wait()
}
