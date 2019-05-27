/***********************************************************************
# File Name: init.go
# Author: TiDao
# mail: songzhaofeng@yahoo.com
# Created Time: 2018-09-10 13:30:46
*********************************************************************/
package main

import(
    "image"
)

func loadIcons(){
    icons = map[string]image.Image{
        "spades.png": loadIcon("spades.png")
        "hearts.png": loadIcon("hears.png")
        "diamonds.png": loadIcon("diamaonds.png")
        "clubs.png":loadIcon("clubs.png")
    }
}

func Icon(name string) image.Image{
    if icons == nil{
        loadIcons()
    }
    return icons[name]
}

func loadIcons(){
    icons = make(map[string]image.Image)
    icons["spades.png"] = loadIcon("spades.png")
    icons["hearts.png"] = loadIcon("hearts.png")
    icons["diamonds.png"] = loadIcon("diamonds.png")
    icons["clubs.png"] = loadIcon("clubs.png")
}



