/***********************************************************************
# File Name: bubblesort_test.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2019-04-29 14:03:20
*********************************************************************/
package bubblesort

import "testing"

func TestBubbleSort1(t *testing.T){
    values := []int{1,2,3,4,5,3,2,1}
    Bubblesort(values)
}
