/***********************************************************************
# File Name: qsort_test.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2019-04-29 18:35:15
*********************************************************************/
package qsort

import "testing"
import "math/rand"
import "time"
import "fmt"

func TestQuicksort(t *testing.T){
    cashi := make([]int ,10000)
    for i:=0; i<10000; i++{
        cashi[i] = rand.Intn(1000)
    }
    start_time := time.Now()
    Quicksort(cashi)
    end_time := time.Since(start_time)
    fmt.Println(cashi)
    fmt.Println(end_time)
}
