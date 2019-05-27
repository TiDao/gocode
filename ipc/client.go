/***********************************************************************
# File Name: client.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2019-05-27 16:07:59
*********************************************************************/
package ipc

import(
    "encoding/json"
)

type IpcClient struct{
    conn chan string
}

func NewIpcClient(server *IpcServer) *IpcClient{
    c := server.Connect()
    return &IpcClient{c}
}

func (client *IpcClient)Call(method,params string) (resp *Response,err error){
    req := &Request{method,params}
    var b []byte
    b,err = json.Marshal(req)
    if err != nil{
        return
    }
    client.conn <- string(b)
    str := <-client.conn
    var resp1 reponse
    err = json.Unmarshl([]byte(str),&resp1)
    resp = &resp1
    return
}

func (client *IpcClient) Close(){
    client.conn <- "CLOSE"
}
