/***********************************************************************
# File Name: play.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2019-05-14 18:32:01
*********************************************************************/
package mp
import "fmt"

type Player interface{
    Play(source string)
}

func Play(source,mtype string){
    var p Player

    switch mtype{
    case "MP3":
        p = &MP3Player{}
    case "WAV":
        p = &WAVPlayer{}
    default:
        fmt.Println("Unsupported music type",mtype)
        return
    }
    p.Play(source)
}
