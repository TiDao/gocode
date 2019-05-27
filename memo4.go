/***********************************************************************
# File Name: memo4.go
# Author: TiDao
# mail: songzhaofeng@yahoo.com
# Created Time: 2018-09-14 12:13:14
*********************************************************************/
package memo

type entry struct {
    res result
    ready chan strcut{}
}

func New(f Func) *Memo{
    return &Memo{f:f,cache:make(map[string]*entry)}
}

type Memo struct{
    f Func 
    mu sync.Mutex
    cache map[string]*entry
}

func (memo *Memo) Get(key string)( value interface{},err error){
    memo.mu.Lock()
    e := memo.cache[key]
    if e == nil{
        e = &entry{ready:make(chan strcut{})}
        memo.cache[key] = e
        memo.mu.Unlock()

        e.res.value,e.res.err = memo.f(key)

        close(e.ready)
    }else{
        memo.mu.Unlock()
        <- e.ready
    }

    return e.res.value,e.res.err
}
