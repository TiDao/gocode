/***********************************************************************
# File Name: coloredpoint.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2019-02-13 11:43:49
*********************************************************************/

package  main
import(
    "image/color"
)

type Point struct{
    X,Y float64
}

type ColoredPoint struct{
    Point
    Color color.RGBA
}


