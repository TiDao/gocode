/***********************************************************************
# File Name: bzipper.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2018-11-30 11:23:39
*********************************************************************/
package main

import(
    "io"
    "log"
    "os"
    "gopl.io/ch13/bzip"
)

func main(){
    w := bzip.NewWriter(os.Stdout)
    if _,err := io.Copy(w,os.Stdin); err != nil{
        log.Fatalf("bzipper:%v\n",err)
    }

    if err := w.Close(); err != nil{
        log.Fatalf("bzipper:close : %v\n",err)
    }
}

