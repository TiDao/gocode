/***********************************************************************
# File Name: channel.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2019-04-02 15:15:41
*********************************************************************/
package main

import "fmt"

func Count (ch chan int){
    fmt.Println("Counting")
    ch <-1
}

func main(){
    chs := make([]chan int,10)
    for i:=0; i<10;i++{
        chs[i] = make(chan int)
        go Count(chs[i])
    }

    for _,ch := range(chs){
        <-ch
    }
}
