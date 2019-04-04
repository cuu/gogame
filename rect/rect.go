package rect

import (
	"github.com/veandco/go-sdl2/sdl"
)

func NewRect(coords ...int) *sdl.Rect {
	r := &sdl.Rect{0,0,0,0}

	for i,c := range(coords) {
		if i==0 {
			r.X = int32(c)
		}
		if i==1 {
		  r.Y = int32(c)
		}
		if i==2 {
			r.W = int32(c)
		}
		if i==3 {
			r.H = int32(c)
		}
	}
	
	return r
}

func Rect(top ,left, width,height int) sdl.Rect {
	return sdl.Rect{int32(top),int32(left),int32(width),int32(height)}
}

func InflateIp(rect *sdl.Rect,  x,y int) {

	rect.X -= int32(x/2)
	rect.Y -= int32(y/2)
	rect.W += int32(x)
	rect.H += int32(y)
	
}

func Inflate(rect *sdl.Rect,  x,y int)  sdl.Rect {

	r := sdl.Rect{0,0,0,0}
	
	r.X = rect.X - int32(x/2)
	r.Y = rect.Y - int32(y/2)
	r.W = rect.W + int32(x)
	r.H = rect.H + int32(y)

	return r
}
