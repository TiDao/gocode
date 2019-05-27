/***********************************************************************
# File Name: change_value.go
# Author: TiDao
# mail: songzhaofeng@yahoo.com
# Created Time: 2018-11-23 10:10:59
*********************************************************************/

package main

import(
    "fmt"
    "reflect"
)

func main(){
    x := 2
    a := reflect.ValueOf(2)
    b := reflect.ValueOf(x)
    c := reflect.ValueOf(&x)

    fmt.Println(a.CanAddr())
    fmt.Println(b.CanAddr())
    fmt.Println(c.CanAddr())

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)

    fmt.Println("testing")
    fmt.Println(3)


    d := reflect.ValueOf(&x).Elem()
    fmt.Println(d.Addr())
    fmt.Println(d.Addr().Interface())
    fmt.Println(d.Addr().Interface().(*int))

    d.Set(reflect.ValueOf(4))
    fmt.Println(x)
    fmt.Println(reflect.TypeOf(d))
    fmt.Println(d.Kind())

    d.Set(reflect.ValueOf(int64(5)))
}
