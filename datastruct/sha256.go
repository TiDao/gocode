/***********************************************************************
# File Name: sha256.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2018-11-30 14:56:14
*********************************************************************/
package main

import(
    "fmt"
    "crypto/sha256"
)

func main(){
    c1 := sha256.Sum256([]byte("x"))
    c2 := sha256.Sum256([]byte("X"))
    fmt.Printf("%x\n%x\n%t\n%T\n",c1,c2,c1==c2,c1)
    
}
