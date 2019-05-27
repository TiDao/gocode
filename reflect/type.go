/***********************************************************************
# File Name: type.go
# Author: TiDao
# mail: songzhaofeng@yahoo.com
# Created Time: 2018-11-20 11:17:16
*********************************************************************/

package main

import(
    "reflect"
    "fmt"
    "io"
    "os"
)

func main(){
    t := reflect.TypeOf(3)
    fmt.Println(t.String())
    fmt.Println(t)
    var w io.Writer = os.Stdout
    fmt.Println(reflect.TypeOf(w))

    fmt.Printf("%T\n",3)
}
