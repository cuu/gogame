package transform


import (
	"log"
	"errors"
	
	"github.com/veandco/go-sdl2/sdl"
//	"github.com/cuu/gogame/color"
	
)


func newsurf_fromsurf(surf *sdl.Surface, width ,height int) (*sdl.Surface,error) {
	var newsurf *sdl.Surface

	if surf.BytesPerPixel() <= 0 || surf.BytesPerPixel() > 4 {
		log.Fatal("unsupport Surface bit depth for transform")
		return nil,errors.New("unsupport Surface bit depth for transform")
	}

	newsurf,err := sdl.CreateRGBSurface(0,int32(width),int32(height),32,
		surf.Format.Rmask, surf.Format.Gmask, surf.Format.Bmask, surf.Format.Amask)
	if err != nil {
		panic( fmt.Sprintf("newsurf_fromsurf failed %s",sdl.GetError()))
	}

	if surf.BytesPerPixel() == 1 && surf.Format.Palette != nil {
		newsurf.SetPalette(surf.Format.Palette)
	}

	key,err := surf.GetColorKey()
	if err == nil {
		newsurf.SetColorKey(true,key)
	}

	alpha,err := surf.GetAlphaMod()
	if err == nil {
		newsurf.SetAlphaMod(alpha)
	}
	return newsurf
}


func Stretch(src *sdl.Surface, dst *sdl.Surface) {
	srcrow := src.Pixels()
	dstrow := dst.Pixels()
	
	bytes_per_pixel := src.BytesPerPixel()

	srcpitch := src.Pitch
	dstpitch := dst.Pitch
	
	dstwidth := dst.W
	dstheight := dst.H

	dstwidth2 := dst.W << 1
	dstheight2 := dst.H << 1

	srcwidth2 := src.W << 1
	srcheight2 := src.H << 1

	w_err := srcheight2 - dstheight2
	h_err := w_err

	srcrow_addr := 0
	dstrow_addr := 0
	
	switch bytes_per_pixel {
	case 1:
		srcrow_addr = 0
		dstrow_addr = 0
		for looph := 0; looph <dstheight; looph++ {
			srcpix := srcrow_addr
			dstpix := dstrow_addr
			
			w_err = srcwidth2 - dstwidth2
			for loopw :=0; loopw < dstwidth; loopw++ {
				dstrow[dstpix] = srcrow[srcpix]
				dstpix+=1
				for w_err >= 0 {
					srcpix+=1
					w_err -= dstwidth2
				}
				w_err += srcwidth2
			}

			for h_err >= 0 {
				srcrow_addr += srcpitch
				h_err -= dstheight2
			}
			dstrow_addr += dstpitch
			h_err += srcheight2
		}
		break
	case 2:
		srcrow_addr = 0
		dstrow_addr = 0
		for looph := 0; looph < dstheight; looph++ {
			srcpix := srcrow_addr
			dstpix := dstrow_addr
			w_err = srcwidth2 - dstwidth2
			for loopw := 0; loopw < dstwidth; loopw ++ {
				dstrow[dstpix]   = srcrow[srcpix]
				dstrow[dstpix+1] = srcrow[srcpix+1]
				dstpix+=2
				for w_err >= 0 {
					srcpix+=2
					w_err -= dstwidth2
				}
				w_err += srcwidth2
			}
			for h_err >= 0 {
				srcrow_addr += srcpitch
				h_err -= dstheight2
			}
			dstrow_addr += dstpitch
			h_err += srcheight2
		}
		break
	case 3:
		srcrow_addr = 0
		dstrow_addr = 0
		for looph := 0 ; looph < dstheight; looph++ {
			srcpix := srcrow_addr
			dstpix := dstrow_addr
			w_err = srcwidth2 - dstwidth2
			for loopw := 0; loopw <dstwidth; loopw++ {
				dstrow[dstpix]   = srcrow[srcpix]
				dstrow[dstpix+1] = srcrow[srcpix+1]
				dstrow[dstpix+2] = srcrow[srcpix+2]
				dstpix += 3
				for w_err >= 0 {
					srcpix += 3
					w_err -= dstwidth2
				}
				w_err +=srcwidth2
			}

			for h_err >= 0 {
				srcrow_addr += srcpitch
				h_err -= dstheight2
			}

			dstrow_addr += dstpitch
			h_err += srcheight2
		}
		break
	case 4:
		srcrow_addr = 0
		dstrow_addr = 0
		for looph := 0 ; looph < dstheight; looph++ {
			srcpix := srcrow_addr
			dstpix := dstrow_addr
			w_err = srcwidth2 - dstwidth2
			for loopw := 0; loopw <dstwidth; loopw++ {
				dstrow[dstpix]   = srcrow[srcpix]
				dstrow[dstpix+1] = srcrow[srcpix+1]
				dstrow[dstpix+2] = srcrow[srcpix+2]
				dstrow[dstpix+3] = srcrow[srcpix+3]				
				dstpix += 4
				for w_err >= 0 {
					srcpix += 4
					w_err -= dstwidth2
				}
				w_err +=srcwidth2
			}

			for h_err >= 0 {
				srcrow_addr += srcpitch
				h_err -= dstheight2
			}

			dstrow_addr += dstpitch
			h_err += srcheight2
		}
		break		
	}
}


func Scale(src_surf *sdl.Surface, new_width ,new_height int )  *sdl.Surface {
	if src_surf == nil {
		return src_surf
	}

	if new_width < 0 || new_height < 0 {
		panic("Cannot scale to negative size")
	}
	
	newsurf := newsurf_fromsurf(src_surf,new_width,new_height)
	if newsurf.W != new_width || newsurf.H != new_height {
		panic("Destination surface not the given width or height.")
	}

	if src_surf.BytesPerPixel() != newsurf.BytesPerPixel() {
		panic("Source and destination surfaces need the same format.")
	}

	if new_width > 0 && new_height > 0 {
		newsurf.Lock()
		Stretch(surf,newsurf)
		newsurf.Unlock()
	}
	
	return newsurf
}


