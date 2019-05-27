/***********************************************************************
# File Name: equal.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2018-11-28 15:38:02
*********************************************************************/
package main

import(
    "fmt"
    "reflect"
    "strings"
    _ "unsafe"
)

func TestSplit(t *testing.T){
    got := strings.Sprlit("a:b:c",":")
    want := []string{"a","b","c"}
    if !reflect.DeepEqual(got,want){}
}

func main(){
    var a []string =nil
    var b []string ={}
    fmt.Println(reflect.DeepEqual(a,b)) //"false"
    var c,d map[string]int = nil,make(map[string]int)
    fmt.Println(reflect.DeepEqual(c,d))
}

func equal(x,y reflect.Value,seen map[comparison]bool) bool{
    if !x.IsValid() || !y.IsValid(){
        return x.IsValid() == y.IsValid()
    }
    if x.Type() != y.Type(){
        return false
    }

    switch x.Kind(){
    case reflect.Bool:
        return x.Bool() == y.Bool()
    case reflect.String:
        return x.String() == y.String()
    case reflect.Chan,reflect.UnsafePointer,reflect.Func:
        return x.Pointer() == y.Pointer()
    case reflect.Array,reflect.Slice:
        if x.Len() != y.Len(){
            return false
        }
        for i:=0;i<x.Len();i++{
            if !equal(x.Index(i),y.Index(i),seen){
                return false
            }
        }
        return true
    }
    panic("unreachable")
}

func Equal(x,y interface{} ) bool{
    seen := make(map[comparison]bool)
    return equal(reflect.ValueOf(x),reflect.ValueOf(y),seen)
}

type comparison struct{
    x,y unsafe.Pointer
    t reflect.Type
}


