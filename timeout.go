/***********************************************************************
# File Name: timeout.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2019-05-27 11:24:50
*********************************************************************/
package main

import(
    "fmt"
    "time"
)

func main(){
    timeout := make(chan int,1)
    go func(){
        time.Sleep(1e9)
        timeout <- 1
    }()

    select{
    case <- ch:
    }

}
