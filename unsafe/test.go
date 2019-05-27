/***********************************************************************
# File Name: test.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2018-11-28 16:50:44
*********************************************************************/
package main

import(
    "reflect"
    "fmt"
)

func main(){
    t := reflect.TypeOf(3)
    fmt.Println(reflect.TypeOf(t))
}
