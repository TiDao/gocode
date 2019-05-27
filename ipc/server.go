/***********************************************************************
# File Name: server.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2019-05-27 15:38:29
*********************************************************************/
package ipc

import(
    "encoding/json"
    "fmt"
)

type Request struct{
    Method string "method"
    Params string "params"
}

type Response struct{
    Code string "code"
    Body string "body"
}

type server interface{
    Name() string
    Handle(method,params string) *Response
}

type IpcServer struct{
    Server
}

func (server *IpcServer)Connect() chan string{
    session := make(chan string,0)
    go func(c chan string){
        for{
            request := <-c
            if request == "CLOSE"{
                break
            }
            var req Request
            err := json.Unmarshal([]byte(request),&req)
            if err != nil{
                fmt.Println("Invalied request format:",request)
            }
            resp := server.Handle(req.Method,req.Params)
            b,err := json.Marshal(resp)
            c <- string(b)
        }
        fmt.Println("Session closed")
    }(session)
    fmt.Println("A new session has been created sucessfully")
    return session
}
