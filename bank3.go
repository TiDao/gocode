/***********************************************************************
# File Name: bank3.go
# Author: TiDao
# mail: songzhaofeng@yahoo.com
# Created Time: 2018-09-07 16:29:23
*********************************************************************/
package bank

import "sync"

var (
    mu sync.Mutex
    balance int
)

func Deposit(amount int){
    mu.Lock()
    balance = balance + amount
    mu.Unlock()
}

func Balance() int{
    mu.Lock()
    b := balance
    mu.Unlock()
    return b
}


