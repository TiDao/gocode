/***********************************************************************
# File Name: select.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2019-05-21 16:04:59
*********************************************************************/
package main

import "fmt"
import "sync"

//func main(){
//    ch := make(chan int)
//    for {
//        go func(){
//            i := <-ch
//            fmt.Println("value received :",i)
//        }()
//        select{
//        case ch <-0:
//        case ch <-1:
//        }
//    }
//}

func main() {
	ch := make(chan int,10)
    var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
        wg.Add(1)
		go func() {
            defer wg.Done()
			fmt.Println("testing deadlock")
			ch <-1
		}()
	}
    wg.Wait()
    for i:=0 ;i<10;i++{
        fmt.Println(<-ch)
    }
}

