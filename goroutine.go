/***********************************************************************
# File Name: goroutine.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2019-05-20 17:53:12
*********************************************************************/
package main

import "fmt"

func Add(x,y int){
    z := x+y
    fmt.Println(z)
}
func main(){
    for i := 0; i <=10;i++{
        go Add(i,i)
    }
}
