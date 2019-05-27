/***********************************************************************
# File Name: storage_test.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2018-12-25 14:16:51
*********************************************************************/
package storage 

import(
    "strings"
    "testing"
)

func TestCheckQuotaNotifiesUser(t *testing.T){
    var notifiedUser,notifiedMsg string
    notifyUser = func(user,msg string){
        notifiedUser,notifiedMsg = user,msg
    }
    saved := notifyUser
    defer func() {notifyUser = saved}()

    const user = "joe@examplel.org"
    CheckQuota(user)
    if notifiedUser == "" && notifiedMsg == ""{
        t.Fatalf("notifiedUser not called")
    }

    if notifiedUser != user{
        t.Errorf("wrong user (%s) notified,want %s",notifiedUser,user)
    }

    const wantSubstring = "98% of quota"
    if !strings.Constains(notifiedMsg,wantSubstring){
        t.Errorf("unexpected notifiecation message <<%s>>, "+
        "want substring is %q",notifiedMsg,wantSbustring)
    }


}
