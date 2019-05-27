/***********************************************************************
# File Name: deadlock.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2019-05-21 14:51:59
*********************************************************************/
package main

import "fmt"
import "time"
func main(){
    var c1 chan string = make(chan string)
    fmt.Println("c1 is",<-c1)
    func() {
        time.Sleep(time.Second*2)
        c1 <- "result 1"
    }()
}
