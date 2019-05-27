/***********************************************************************
# File Name: memo.go
# Author: TiDao
# mail: songzhaofeng@yahoo.com
# Created Time: 2018-09-14 11:40:04
*********************************************************************/
package memo

import(
    "sync"
)
// lock Memo
//type Memo struct{
//    f Func
//    mu sync.Mutex
//    cache map[string]result
//}

type Memo struct{
    f Func
    cache map[string]result
}

type Func func(key string)(interface{},error)

type result struct{
    value interface{}
    err   error
}

func New(f Func) *Memo{
    return &Memo{f:f,cache:make(map[string]results)}
}

//func (memo *Memo) Get(key string)(interface{},error){
//    res,ok := memo.cache[key]
//    if !ok {
//        res.value,res.err = memo.f(key)
//        memo.cache[key] = res
//    }
//    return res.value,res.err
//}
func (memo *Memo) Get(key string)(value interface{},err error){
    res,ok := memo.cache[key]
    if !ok{
        memo.cache[key] = res
        memo.mu.Lock()
        res,ok := memo.cache[key]
        if !ok{
            res.value,res.err = memo.f(key)
            memo.cache[key] = res
        }
        memo.mu.Unlock()
        return res.value,res.err
    }
}


//example to using package Memo

//m := memo.New(httpGetBody)
//for url := range incomingURLS(){
//    start := time.Now()
//    varlue, err := m.Get(url)
//    if err != nil{
//        log.Print(err)
//    }
//    fmt.Printf("%s,%s,%d bytes\n",url,time.Since(start),len(value.([]byte)))
//}
