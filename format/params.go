/***********************************************************************
# File Name: params.go
# Author: TiDao
# mail: songzhaofeng@yahoo.com
# Created Time: 2018-11-27 14:04:55
*********************************************************************/
package params

func Unpack(req *http.Request,ptr interface{}) error{
    if err := req.ParseForm(); err != nil{
        return err
    }

    fields = make(map[string]reflect.Value)

    v := reflect.ValueOf(ptr).Elem()
    for i := 0; i<v.NumField();i++{
        fieldInfo := v.Type().Field(i)
        tag := fieldInfo.Tag
        name := tag.Get("http")
        if name == ""{
            name = strings.ToLower(fieldInfo.Name)
        }
        field[name] = v.Field(i)
    }

    // update struct field for each parameter in the request.
    for name,values := range req.Form{
        f := field[name]
        if !f.IsValid(){
            continue
        }
        for _,value := range values{
            if f.Kind() == reflect.Slice{
                elem := reflect.New(f.Type().Elem()).Elem()
                if err := populate(elem,value); err != nil{
                    return fmt.Errorf("%s : %v",name,err)
                }
                f.Set(reflect.Append(f,elem))
            }else {
                if  err := populate(f,value);err !=nil{
                    return fmt.Errorf("%s: %v",name,err)
                }
            }
        }
    }
    return nil
}

func populate(v reflect.Value,value string) error{
    switch v.Kind(){
    case reflect.String:
        v.SetString(value)
    case reflect.Int:
        i,err := strconv.ParseInt(value,10,64)
        if err != nil{
            return err
        }
        v.SetInt(i)
    case reflect.Bool:
        b,err := strconv.ParseBool(value)
        if err != nil{
            return err
        }
        v.SetBool(b)
    default:
        return fmt.Errorf("unsupported kind %s",v.Type())
    }
    return nil
}

func Print(x interface{}){
    v := reflect.ValueOf(x)
    t := v.Type()
    fmt.Printf("type %s\n",t)

    for i:=0;i<v.NumMethod();i++{
        methType := v.Method(i).Type()
        fmt.Printf("func (%s) %s%s\n",t,t.Method(i).Name,strings.TrimPrefix(methType.String(),"func"))
    }
}
