/***********************************************************************
# File Name: back.go
# Author: TiDao
# mail: songzhaofeng@yahoo.com
# Created Time: 2018-09-06 16:07:40
*********************************************************************/
package bank 

var deposits = make(chan int)
var balances = make(chan int)

func Deposit(amount int){deposits <- amount}
func Balance() int { return <- balances}

func teller(){
    var balance int
    for {
        select{
        case amount := <-deposits:
            balance += amount 
        case balances <- balances:
        }
    }
}


func init(){
    go teller()
}

type Cake struct{state string}

func baker(cooked chan<-*Cake){
    for {
        cake := new(Cake)
        cake.state = "cooked"
        cooked <- cake
    }
}

func icer(iced chan<- *Cake,cooked <-chan *Cake){
    for cake := range cooked{
        cake.state = "iced"
        iced <-cake
    }
}


