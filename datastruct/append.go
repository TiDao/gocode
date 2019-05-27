/***********************************************************************
# File Name: append.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2018-11-30 15:12:54
*********************************************************************/
package main

import(
    "fmt"
)

func main(){
    var runes []rune
    for _,r := range "hello,世界"{
        runes = append(runes,r)
    }
    fmt.Println(runes)
    fmt.Printf("%q\n",runes)

  
}

func appendInt( x []int,y int) []int{
    var z []int
    zlen := len(x) + 1
    if zlen <= cap(x){
        z = x[:zlen]
    }else{
        zcap := zlen
        if zcap < 2*len(x){
            zcap = 2*len(x)
        }
        z = make([]int,zlen,zcap)
        copy(z,x)
    }
    z[len(x)] = y
    return z
}


//func testappend(x []int,y int) []int{
//    var z []int
//    zlen := len(x)+1
//    if zlen <= cap(x){
//        z = x[:zlen]
//    //    x = append(x,y)
//    }else{
//        zcap := zlen
//        if zcap < 2*len(x){
//            zcap = 2*len(x)
//        }
//        z = make([]int,zlen,zcap)
//        copy(x,y)
//    }
//    z[len(x)] = y
//    return z
//}
