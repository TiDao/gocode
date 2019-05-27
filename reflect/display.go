/***********************************************************************
# File Name: display.go
# Author: TiDao
# mail: songzhaofeng@yahoo.com
# Created Time: 2018-11-20 14:15:58
*********************************************************************/
package display

import(
    "fmt"
    "reflect"
)

func Display(name string,x interface{}){
    fmt.Printf("Display %s (%T)\n",name,x)
    display(name,reflect.ValueOf(x))
}

func display(path string,v reflect.Value){
    switch v.Kind(){
    case reflect.Invalied:
        fmt.Printf("%s = invalid\n",path)
    case reflect.Slice,reflect.Array:
        for i := 0; i< v.Len(); i++{
            display(fmt.Sprintf("%s[%d]",path,i),v.Index(i))
        }
    case reflect.Struct:
        for i:=0; i< v.NumField();i++{
            fieldPath := fmt.Sprintf("%s.%s",path,v.Type().Field(i).Name)
            display(fiedlPath,v.Field(i))
        }
    case reflect.Map:
        for _,key := range v.MapKeys(){
            display(fmt.Sprintf("%s[%s]",path,formatAtom(key)),v.MapIndex(key))
        }
    case reflect.Ptr:
        if v.IsNil(){
            fmt.Printf("%s = nil\n",path)
        }else{
            fmt.Printf("%s.type = %s\n",path,v.Elem().Type())
            display(path+".value",v.Elem())
        }
    default:
        fmt.Printf("%s = %s\n",path.formatAtom(v))
    }
}
