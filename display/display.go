package display

import (
	"os"
	
	"github.com/veandco/go-sdl2/sdl"
	"github.com/cuu/gogame"
	"github.com/cuu/gogame/surface"
	
)

var Inited =  false
var window *sdl.Window
var win_surface *sdl.Surface
var big_surface *sdl.Surface

func AssertInited() {
	if Inited == false {
		panic("run gogame.DisplayInit first")
	}
}

func Init() bool {
	sdl.Do(func() {
		
		if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
			panic(err)
		}
	
		Inited = true
	})
	
  return Inited 
}


func SetMode(w,h,flags,depth int32) *sdl.Surface {
	var err error
	var surf *sdl.Surface
	AssertInited()
	
	sdl.Do(func() {
		video_centered := os.Getenv("SDL_VIDEO_CENTERED")
		if video_centered == "1" {
			window, err = sdl.CreateWindow("gogame", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
				w, h, uint32( gogame.SHOWN | flags))
		}else {
			window, err = sdl.CreateWindow("gogame", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
				w, h, uint32( gogame.SHOWN | flags))
		}
		
		if err != nil {
			panic(err)
		}

		surf,err = window.GetSurface()
		if err != nil {
			panic(err)
		}
		
		win_surface = surf
		big_surface = surface.Surface(int(win_surface.W),int(win_surface.H))

		
	})

	return big_surface
}

func Flip(  ) {
	sdl.Do(func() {
		if win_surface != nil && big_surface != nil {
			surface.Blit(win_surface,big_surface, nil,nil)
		}
		if window != nil {
			window.UpdateSurface()
		}
	})
}
		


