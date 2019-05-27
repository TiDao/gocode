/***********************************************************************
# File Name: bubblesort.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2019-04-29 13:40:58
*********************************************************************/
package bubblesort

import "fmt"

//func Bubblesort(values []int){
//    for i := 0;i < len(values) -1; i++{
//        for j := i+1; j < len(values);j++{
//            if values[i] > values[j]{
//                values[i],values[j] = values[j],values[i]
//            }
//        }
//    }
//    fmt.Println(values)
//}

func Bubblesort(values []int){
    flag := true
    for i := 0 ;i <len(values)-1;i++{
        flag = true
        fmt.Println(i)
        for j := 0; j < len(values)-i-1;j++{
            fmt.Println(j)
            if values[j] > values[j+1]{
                values[j],values[j+1] = values[j+1],values[j]
                fmt.Println("test")
                flag = false
            }
        }
        if flag == true{
            fmt.Println(values)
            break
        }
    }
}
