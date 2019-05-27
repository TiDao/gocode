/***********************************************************************
# File Name: search.go
# Author: TiDao
# mail: songzhaofeng@yahoo.com
# Created Time: 2018-11-27 11:52:23
*********************************************************************/

package main

import(
    "gopl.io/ch12/parames"
    "fmt"
    "net/http"
)

func  search(resp http.ResponseWriter,req *http.Request){
   var data struct{
       Labelsd []string `http:"l"`
       MaxResults int    `http:"max"`
       Exact   bool  `http:"x"`
   }
   data.MaxResults = 10
   if err := params.Unpack(req,&data); err!=nil{
       http.Error(resp,err.Error(),http.StatusBadRequest)
       return
   }
   fmt.Fprintf(resq,"Search %+v\n",,data)
}


