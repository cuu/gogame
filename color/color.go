package color

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Color struct {
	R uint32
	G uint32
	B uint32
	A uint32
}

func NewColorSDL(color sdl.Color) *Color {
	_color := &Color{0,0,0,255}

  _color.R = uint32(color.R)
  _color.G = uint32(color.G)
  _color.B = uint32(color.B)
  _color.A = uint32(color.A)
	
	return _color
}

func NewColor(colors ...int) *Color {
	color := &Color{0,0,0,255}

	for i,c := range(colors) {
		if i==0 {
			color.R = uint32(c)
		}
		if i==1 {
		  color.G = uint32(c)
		}
		if i==2 {
			color.B = uint32(c)
		}
		if i==3 {
			color.A = uint32(c)
		}
	}
	
	return color
}

func (c *Color) ToSDL() sdl.Color {
	sdl_color := sdl.Color{0,0,0,0}

	sdl_color.R = uint8(c.R)
	sdl_color.G = uint8(c.G)
	sdl_color.B = uint8(c.B)
	sdl_color.A = uint8(c.A)
	return sdl_color
}

func (c *Color) ToHex() int {
	return int( c.A<< 24 | c.R << 16 | c.G << 8 | c.B )
}

func (c *Color) ToBytes() []byte {
	bytes := make([]byte,4)
	bytes[0] = byte(c.R)
	bytes[1] = byte(c.G)
	bytes[2] = byte(c.B)
	bytes[3] = byte(c.A)
	return bytes
}

func (c *Color) RGBA() (r, g, b, a uint32) {
  r = uint32(c.R)
  r |= r << 8
  g = uint32(c.G)
  g |= g << 8
  b = uint32(c.B)
  b |= b << 8
  a = uint32(c.A)
  a |= a << 8
  return
}
