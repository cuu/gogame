package time

import (
//"sync"
//	"fmt"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/cuu/gogame/event"
//	"time"
)


func Delay( dur int ) {

	go func() {
		event.Pause()		
		sdl.Delay( uint32(dur))
		event.Resume()

	}()
	
}
