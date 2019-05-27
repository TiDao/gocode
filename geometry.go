/***********************************************************************
# File Name: geometry.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2019-02-13 10:21:17
*********************************************************************/
package geometry

import "math"

type Point struct{ X,y float64}

func Distance(p,q Point) float64{
    return math.Hypot(q.X-p,q.Y-p.Y)
}

func (p Point) Distance(q Point) float64{
    return math.Hypot(q.X-p.X,q.Y-p.Y)
}


type IntList struct{
    Value int
    Tail *IntList
}

func (list *IntList) Sum() int{
    if list == nil{
        return 0
    }
    return list.Value + list.Tail.Sum()
}


