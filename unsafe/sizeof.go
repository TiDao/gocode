/***********************************************************************
# File Name: sizeof.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2018-11-27 18:05:32
*********************************************************************/
package main

import(
    "unsafe"
    "fmt"
)

func main(){
    fmt.Println(unsafe.Sizeof(float64(0)))
    fmt.Println(unsafe.Sizeof(int8(0)))
    fmt.Println(unsafe.Sizeof(interface{}(0)))
    fmt.Println(unsafe.Sizeof(string("a")))
    fmt.Println(unsafe.Sizeof([]string{"a"}))
}
