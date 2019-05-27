/***********************************************************************
# File Name: bank2.go
# Author: TiDao
# mail: songzhaofeng@yahoo.com
# Created Time: 2018-09-06 17:45:24
*********************************************************************/
package main
var (
    sema = make(chan struct{},1)
    balance int
)

func Deposit(amount int){
    sema <- struct{}
    balance = balance + amount
    <- sema
}

func Balance() int{
    sema <- struct{}
    b := balance
    <-sema
    return b
}
