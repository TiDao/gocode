/***********************************************************************
# File Name: test.go
# Author: TiDao
# mail: songzhaofeng@yahoo.com
# Created Time: 2018-09-06 17:48:40
*********************************************************************/
package main

import "fmt"


func first( ch chan<- *int){
        num := new(int)
        *num = 1
        ch <- num
}

func second(iced chan<- *int,ch <-chan *int){
    for num := range ch{
        *num = 2
        iced <- num
    }
}

func main(){
    chs := make(chan *int)
    chout := make(chan *int)
    for i:=0;i<5;i++{
        go second(chout,chs)
        go first(chs)
        fmt.Println(*<-chout)
    }
}
