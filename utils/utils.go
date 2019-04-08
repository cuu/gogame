package utils

import (
	"log"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/cuu/gogame/color"
)


func ColorSurface(surf *sdl.Surface,col *color.Color) {

	bytes_per_pixel := surf.BytesPerPixel()
	if bytes_per_pixel <= 0 || bytes_per_pixel > 4 {
		log.Fatalf("unsupport surface format %d\n",bytes_per_pixel)
		return
	}
	
	color_bytes := col.ToBytes()
	
	surf.Lock()
	pixels := surf.Pixels()

	for i:= 0; i< surf.PixelNum(); i++ {
		for j:=0;j< bytes_per_pixel;j++ {
			pixels[i*bytes_per_pixel+j] = color_bytes[j]
		}
	}
	
	surf.Unlock()
}


