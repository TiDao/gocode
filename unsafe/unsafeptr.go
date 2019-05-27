/***********************************************************************
# File Name: unsafeptr.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2018-11-28 14:00:48
*********************************************************************/
package main

import(
    "fmt"
    "unsafe"
)

func main(){
    var x struct{
        a bool
        b int16
        c []int
    }

    pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
    *pb = 42
    fmt.Println(x.b)
}

