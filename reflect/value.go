/***********************************************************************
# File Name: value.go
# Author: TiDao
# mail: songzhaofeng@yahoo.com
# Created Time: 2018-11-20 11:26:42
*********************************************************************/

package main

import(
    "fmt"
    "reflect"
)

func main(){
    v := reflect.ValueOf(3)
    fmt.Println(v)
    fmt.Printf("%v\n",v)
    fmt.Printf("%T\n",v)
    fmt.Println(v.String())
    t := reflect.TypeOf(v)
    fmt.Println(t)

    m := v.Type()
    fmt.Println(m)  // print the value of m , is int
    fmt.Printf("%T\n",m) // print the type of m ,is  *reflect.rtype
    fmt.Printf("%v\n",m) // print the value of m, is int
    fmt.Printf("%v\n",m.String()) // print the return value of m.String ,is int; the m.String() type is string;
}
