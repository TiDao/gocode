/***********************************************************************
# File Name: align.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2018-11-27 18:25:38
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
    fmt.Printf("x size is %d,x align is %d\n",unsafe.Sizeof(x),unsafe.Alignof(x))
    fmt.Printf("a size is %d,a align is %d,a offset is %d\n",
    unsafe.Sizeof(x.a),unsafe.Alignof(x.a),unsafe.Offsetof(x.a))
    fmt.Printf("b size is %d,b align is %d,a offset is %d\n",
    unsafe.Sizeof(x.b),unsafe.Alignof(x.b),unsafe.Offsetof(x.b))
    fmt.Printf("c size if %d,c align is %d,c offset is %d\n",
    unsafe.Sizeof(x.c),unsafe.Alignof(x.c),unsafe.Offsetof(x.c))
}
