package time

import (
//"sync"
//	"fmt"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/cuu/gogame/event"
	gotime "time"
)


func Delay( dur int ) {

	go func() {
		event.Pause()		
		sdl.Delay( uint32(dur))
		event.Resume()

	}()
	
}

func BlockDelay( dur int ) {

	event.Pause()		
	sdl.Delay( uint32(dur))
	event.Resume()

}

func Unix() int32 {
    return int32(gotime.Now().Unix())
}
