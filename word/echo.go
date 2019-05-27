/***********************************************************************
# File Name: echo.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2018-12-25 10:16:42
*********************************************************************/
package main

import(
    "flag"
    "fmt"
    "io"
    "os"
    "strings"
)

var (
    n = flag.Bool("n",false,"omint trailing newline")
    s = flag.String("s"," ","separator")
)
var out io.Writer = os.Stdout

func main(){
    flag.Parse()
    if err := echo(!*n,*s,flag.Args());err != nil{
        fmt.Fprintf(os.Stderr,"echo:%v\n",err)
        os.Exit(1)
    }
}

func echo(newline bool,sep string,args []string) error{
    fmt.Fprintf(out,strings.Join(args,sep))
    if newline {
        fmt.Fprintln(out)
    }
    return nil
}
