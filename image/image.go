package image

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/img"

)

func Load(filename string) *sdl.Surface {

	surf,err := img.Load(filename)
	if err != nil {
		panic(fmt.Sprintf("Load %s failed: %s",filename,img.GetError()) )
	}
	
	return surf
}


