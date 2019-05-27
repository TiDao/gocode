/***********************************************************************
# File Name: pointer.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2018-11-28 10:21:55
*********************************************************************/
package main

import(
    "fmt"
    "unsafe"
)

func bits(f float64) uint64{ return *(*uint64)(unsafe.Pointer(&f))}

func main(){
    fmt.Printf("%#016x\n",bits(1.0))
}
