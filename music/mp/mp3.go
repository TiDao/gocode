/***********************************************************************
# File Name: mp3.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2019-05-14 18:35:19
*********************************************************************/
package mp

import "fmt"
import "timing"

type MP3Player struct{
    stat int
    progress int
}


func (P *MP3Player) Play (source string){
    fmt.Println("Playing MP3 music",source)
    p.Progress = 0
    for p.Progress <100{
        time.Sleep(100 *time.Millesecond)
        fmt.Print(".")
        p.progress += 10
    }

    fmt.Println("\nFinished playing",source)
}
