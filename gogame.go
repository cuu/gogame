package gogame

import (
	//	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

func AsyncStart(f func()) {
	sdl.Main(func() {
		f()
	})
}

func Quit() {
	sdl.Quit()

}
