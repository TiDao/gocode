/***********************************************************************
# File Name: array.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2018-11-30 11:54:52
*********************************************************************/
package main

import(
    "fmt"
)

func main(){
    var a [3]int
    fmt.Println(a)
    fmt.Println(len(a)-1)
    for i,v := range a{
        fmt.Printf("%d %d\n",i,v)
    }

    for _ ,v := range a{
        fmt.Printf("%d \n",v)
    }

    var q [3]int = [3]int{1,2,3}
    var r [3]int = [3]int{1,2}
    fmt.Println(q[0])
    fmt.Println(r[2])

    f := [...]int{1,2,3}
    fmt.Printf("%T\n",f)

}

func zero(ptr i*[32]byte){
    for i := range ptr{
        ptr[i] = 0
    //  ptr[i] = [32]byte{}
    }
}
