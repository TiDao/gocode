/***********************************************************************
# File Name: once.go
# Author: TiDao
# mail: songzhaofeng@yahoo.com
# Created Time: 2018-09-10 17:42:22
*********************************************************************/
package main

import image
import sync

var icons map[string]image.Image

func loadIcons(){
    icons = map[string]image.Image{
        "spades.png":loadIcon("spades.png")
        "hearts.png":loadIcon("hearts.png")
        "diamonds.png":loadIcon("diamonds.png")
        "clubs.png": loadIcon("clubs.png")
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

var mu sync.Mutex
var icons map[string]image.Image

func Icon(name string) image.Image{
    mu.Lock()
    defer mu.Unlock()
    if icons == nil{
        loadIcons()
    }
    return icons[name]
}

var mu sync.RWMutex
var icons map[string]image.Image

func Icon(name string) image.Image{
    mu.RLock()
    if icons != nil{
        icon := icons[name]
        mu.RUnlock()
        return icon
    }
    mu.RUnlock()

    mu.Lock()
    if icons == nil{
        loadIcons()
    }
    icon := icons[name]
    mu.Unlock()
    return icon
}



var loadIconOnce sync.Once
var icons map[string]image.Image 
func Icon(name string) image.Image{
    loadIconsOj
}
