/***********************************************************************
# File Name: gomaxprocs.go
# Author: TiDao
# mail: songzhaofeng@yahoo.com
# Created Time: 2018-09-26 15:25:18
*********************************************************************/
package main

import(
    "fmt"
)

func main(){
    for{
        go fmt.Print(0)
        fmt.Print(1)
    }
}
