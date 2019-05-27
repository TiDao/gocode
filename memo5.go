/***********************************************************************
# File Name: memo5.go
# Author: TiDao
# mail: songzhaofeng@yahoo.com
# Created Time: 2018-09-23 15:16:14
*********************************************************************/
package main

type request struct{
    key string
    response chan<- result
}

type Memo struct{ requests chan request}

func New(f Func) *Memo{
    memo := &Memo{requests: make(chan request)}
    go memo.server(f)
    return memo
}


func (memo *Memo) Get(key string) (interface{},error){
    reponse := make(chan result)
    memo.requset<- requsest{key,response}
    res := <-reponse
    return res.value, res.err
}

func (memo *Memo) close(){close(memo.requests)}


