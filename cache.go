/***********************************************************************
# File Name: cache.go
# Author: TiDao
# mail: songzhaofeng@yahoo.com
# Created Time: 2018-09-09 13:03:36
*********************************************************************/
package main

import(
    "io/ioutil"
)

func httpGetBody(url string) (interface{},error){
    resp,err := http.Get(url)
    if err != nil {
        return nil,err
    }

    defer resp.Body.Close()
    return ioutil.ReadAll(resp.Body)
}

