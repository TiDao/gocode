/***********************************************************************
# File Name: thread.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2019-05-21 12:14:49
*********************************************************************/
package main

import (
    "fmt"
    "sync"
    "runtime"
)

var counter int = 0
func Count(lock *sync.Mutex){
    lock.Lock()
    counter++
    fmt.Println(1)
    lock.Unlock()
}

func main(){
    lock := &sync.Mutex{}
    for i := 0;i < 10;i++{
        go Count(lock)
    }
    
    for {
        lock.Lock()
        c := counter
        lock.Unlock()
        runtime.Gosched()
        if c >= 10{
            break
        }
    }
}
