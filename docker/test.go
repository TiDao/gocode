/***********************************************************************
# File Name: test.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2018-12-06 15:34:14
*********************************************************************/
package main

import(
    "fmt"
    "os"
)

func main(){
    fmt.Println(os.Readlink("/proc/self/exe"))
}
