/***********************************************************************
# File Name: readlink.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2018-12-11 13:11:22
*********************************************************************/
package main

import(
    "fmt"
    "os"
    "os/exec"
    "log"
    "bufio"
)

func main(){
    cmd,err := os.Readlink("/proc/self/exe")
    if err != nil{
        log.Fatal(err)
    }

    fmt.Println(cmd)

    result := exec.Command("/proc/self/exe")
    fmt.Println(result.Args)
    stdout,erro := result.StdoutPipe()
    if erro != nil{
        log.Fatal(erro)
    }

    result.Start()
    reader := bufio.NewReader(stdout)
    line ,erro2 := reader.ReadString('\n')
    if erro2 != nil{
        log.Fatal(erro2)
    }
    fmt.Printf(line)
    result.Wait()
}
