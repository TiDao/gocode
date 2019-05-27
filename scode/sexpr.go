/***********************************************************************
# File Name: sexpr.go
# Author: TiDao
# mail: songzhaofeng@yahoo.com
# Created Time: 2018-11-22 15:33:44
*********************************************************************/
package main

import(
    "fmt"
    "reflect"
    "bytes"
)

func encode(buf *bytes.Buffer,v reflect.Value) error{
    switch v.Kind(){
    case reflect.Invalid:
        buf.writerString("nil")
    case reflect.Int,reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64:
        fmt.Fprintf(buf,"%d",v.Int())
    case reflect.Uint,reflect.Uint8,reflect.Uint16,reflect.Uint32,reflect.Uint64,reflect.Uintptr:
        fmt.Fprintf(buf,"%d",v.Uint())
    case reflect.String:
        fmt.Fprintf(buf,"%q",v.String())
    case reflect.Ptr:
        return encode(buf,v.Elem())
    case reflect.Array,reflect.Slice:
        buf.WriteByt()
    }
}
