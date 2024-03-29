package time

import (
	//"sync"
	//"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"math"

	"github.com/cuu/gogame/event"
	gotime "time"
)

type Clock struct {
	last_time gotime.Time
}

var myclock *Clock

func NewClock() *Clock {
	if myclock == nil {
		myclock = &Clock{}
	}
	return myclock
}

func (self *Clock) Tick(framerate ...int) int {
	_framerate := 30
	if len(framerate) > 0 {
		_framerate = framerate[0]
	}

	speed := 1000.0 / float64(_framerate)
	//speed = speed * 1000.0

	now := gotime.Now()

	if self.last_time.IsZero() {
		self.last_time = now
		return int(speed)
	} else {
		delta_ms := int(now.Sub(self.last_time).Nanoseconds() / 1e6)

		if delta_ms < int(speed) {
			//fmt.Println("block delayed",delta_ms,speed)
			SDL_Delay(int(math.Floor(speed)))
		} else {

			//fmt.Println("No block delayed",delta_ms,speed)
		}
		self.last_time = now
	}

	return int(speed)
}

func Delay(dur int) {

	go func() {
		event.Pause(sdl.KEYDOWN)
		sdl.Delay(uint32(dur))
		event.Resume(sdl.KEYDOWN)

	}()

}

func SDL_Delay(dur int) {
	sdl.Delay(uint32(dur))
}

func BlockDelay(dur int) {

	event.Pause(sdl.KEYDOWN)
	sdl.Delay(uint32(dur))
	event.Resume(sdl.KEYDOWN)

}

func Unix() int32 {
	return int32(gotime.Now().Unix())
}
