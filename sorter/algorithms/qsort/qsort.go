/***********************************************************************
# File Name: qsort.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2019-04-29 15:45:23
*********************************************************************/
package qsort

//import "sync"

func quickSort(values []int,left,right int){
    temp := values[left]
    p := left
    i,j := left,right
    //var wg sync.WaitGroup
    for i <= j{
        for j >=p && values[j] >= temp{
            j--
        }
        if j >= p{
            values[p] = values[j]
            p = j
        }
        for values[i] < temp && i <=p {
            i++
        }
        if i <= p{
            values[p]= values[i]
            p = i
        }
    }
    values[p] = temp
    if p - left > 1{
        go quickSort(values,left,p-1)
    }
    if right -p > 1{
        go quickSort(values,p+1,right)
    }
}

func Quicksort(values []int){
    quickSort(values,0,len(values)-1)
}
