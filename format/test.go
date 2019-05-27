/***********************************************************************
# File Name: test.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2018-11-27 16:13:57
*********************************************************************/
package main

import(
    "fmt"
    "reflect"
    "time"
    "strings"
)

func Print(x interface{}){
    v := reflect.ValueOf(x)
    fmt.Println(reflect.TypeOf(v))
    t := v.Type()
    fmt.Printf("type %s\n",t)

    for i:=0;i<v.NumMethod();i++{
        methType := v.Method(i).Type()
        fmt.Println(methType)
        fmt.Println(methType.String())
        fmt.Println(t.Method(i).Name)
        fmt.Println(strings.TrimPrefix(methType.String(),"func"))
        fmt.Printf("func (%s) %s%s\n",t,t.Method(i).Name,strings.TrimPrefix(methType.String(),"func"))
    }
}


func main(){
    var a int = 1
    Print(a)
    var b string = "test"
    Print(b)
    var c time.Duration
    Print(c)
}
