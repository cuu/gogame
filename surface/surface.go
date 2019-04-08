package surface

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/cuu/gogame/color"
)


func GetWidth(s *sdl.Surface) int {
	return int(s.W)
}

func GetHeight(s *sdl.Surface) int {
	return int(s.H)
}


func Fill(surface *sdl.Surface, col *color.Color) {

	rect := sdl.Rect{0,0,0,0}

	rect.W = surface.W
	rect.H = surface.H

	FillRect(surface, &rect, uint32(col.ToHex()))
}

func FillRect(surface *sdl.Surface,rect *sdl.Rect, color uint32) {
	
	sdl.Do(func() {
		surface.FillRect(rect,color)
	})
	
	return
}

// Create a New Surface
func Surface(w,h int) *sdl.Surface {
	//flags=0, depth=0, masks=None
	Rmask := 0x000000ff
	Gmask := 0x0000ff00
	Bmask := 0x00ff0000
	Amask := uint32(0xff000000)
	
	flags := 0
	depth := 32
	
	surf,err := sdl.CreateRGBSurface(uint32(flags),int32(w),int32(h), int32(depth), uint32(Rmask), uint32(Gmask), uint32(Bmask), uint32(Amask))
	if err != nil {
		panic( fmt.Sprintf("sdl.CreateRGBSurface failed %s",sdl.GetError()))
	}

	return surf
}
// Create a New Surface with more arguments
func ASurface(w,h int,flags uint32, depth int32) *sdl.Surface {
	//flags=0, depth=0, masks=None
	/*
	Bit    7  6  5  4  3  2  1  0
	Data   R  R  R  G  G  G  B  B
	*/
	//surf->PixelFormat->BitsPerPixel=8
	Rmask := 0x0
	Gmask := 0x0
	Bmask := 0x0
	Amask := 0x0
		
	surf,err := sdl.CreateRGBSurface(uint32(flags),int32(w),int32(h), int32(depth), uint32(Rmask), uint32(Gmask), uint32(Bmask), uint32(Amask))
	if err != nil {
		panic( fmt.Sprintf("sdl.CreateRGBSurface 8bit failed %s",sdl.GetError()))
	}

	return surf
}

//dest represents coord on dst surface, shows where you want  put the source surface on dst 
func Blit(dst *sdl.Surface, source *sdl.Surface, dest *sdl.Rect,area *sdl.Rect) sdl.Rect {
	//area == nil copy the entire source to dst
	//dest should not be nil
	dx := 0
	dy := 0
	if dest == nil {
		dx = 0
		dy = 0
	}else{
		dx = int(dest.X)
		dy = int(dest.Y)
	}
	
	dest_rect := sdl.Rect{0,0,0,0}
	dest_rect.X = int32(dx)
	dest_rect.Y = int32(dy)

	dest_rect.W = source.W
	dest_rect.H = source.H

	if area != nil {
		dest_rect.W = area.W
		dest_rect.H = area.H
	}
	
	err := source.Blit( area, dst, &dest_rect)
	if err !=nil {
		return sdl.Rect{0,0,0,0}
	}
	return dest_rect
}
