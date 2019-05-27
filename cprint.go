/***********************************************************************
# File Name: cprint.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2019-04-02 15:37:01
*********************************************************************/
package main

/*
#include <stdio.h>
*/

import "C"
import "unsafe"

func main(){
    cstr := C.CString("hello, world")
    C.Puts(cstr)
    C.Pree(unsafe.Pointer(cstr))
}
