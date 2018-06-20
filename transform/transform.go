package transform


import (
	"fmt"
	"log"
	"errors"
	"encoding/binary"
	
	"github.com/veandco/go-sdl2/sdl"
//	"github.com/cuu/gogame/color"
	
)

//var Endian binary.LittleEndian
var  Uint16Converter =  binary.LittleEndian.Uint16

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
	return newsurf,nil
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
		for looph := 0; looph < int(dstheight); looph++ {
			srcpix := srcrow_addr
			dstpix := dstrow_addr
			
			w_err = srcwidth2 - dstwidth2
			for loopw :=0; loopw < int(dstwidth); loopw++ {
				dstrow[dstpix] = srcrow[srcpix]
				dstpix+=1
				for w_err >= 0 {
					srcpix+=1
					w_err -= dstwidth2
				}
				w_err += srcwidth2
			}

			for h_err >= 0 {
				srcrow_addr += int(srcpitch)
				h_err -= dstheight2
			}
			dstrow_addr += int(dstpitch)
			h_err += srcheight2
		}
		break
	case 2:
		srcrow_addr = 0
		dstrow_addr = 0
		for looph := 0; looph < int(dstheight); looph++ {
			srcpix := srcrow_addr
			dstpix := dstrow_addr
			w_err = srcwidth2 - dstwidth2
			for loopw := 0; loopw < int(dstwidth); loopw ++ {
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
				srcrow_addr += int(srcpitch)
				h_err -= dstheight2
			}
			dstrow_addr += int(dstpitch)
			h_err += srcheight2
		}
		break
	case 3:
		srcrow_addr = 0
		dstrow_addr = 0
		for looph := 0 ; looph < int(dstheight); looph++ {
			srcpix := srcrow_addr
			dstpix := dstrow_addr
			w_err = srcwidth2 - dstwidth2
			for loopw := 0; loopw <int(dstwidth); loopw++ {
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
				srcrow_addr += int(srcpitch)
				h_err -= dstheight2
			}

			dstrow_addr += int(dstpitch)
			h_err += srcheight2
		}
		break
	case 4:
		srcrow_addr = 0
		dstrow_addr = 0
		for looph := 0 ; looph < int(dstheight); looph++ {
			srcpix := srcrow_addr
			dstpix := dstrow_addr
			w_err = srcwidth2 - dstwidth2
			for loopw := 0; loopw <int(dstwidth); loopw++ {
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
				srcrow_addr += int(srcpitch)
				h_err -= dstheight2
			}

			dstrow_addr += int(dstpitch)
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
	
	newsurf,_ := newsurf_fromsurf(src_surf,new_width,new_height)
	if int(newsurf.W) != new_width || int(newsurf.H) != new_height {
		panic("Destination surface not the given width or height.")
	}

	if src_surf.BytesPerPixel() != newsurf.BytesPerPixel() {
		panic("Source and destination surfaces need the same format.")
	}

	if new_width > 0 && new_height > 0 {
		newsurf.Lock()
		src_surf.Lock()
		Stretch(src_surf,newsurf)
		src_surf.Unlock()
		newsurf.Unlock()
	}
	
	return newsurf
}


func filter_shrink_X_ONLYC(srcpix []byte, dstpix []byte, height,srcpitch,dstpitch, srcwidth,dstwidth int  ) {
	
	srcdiff := srcpitch - (srcwidth * 4 )
	dstdiff := dstpitch - (dstwidth * 4 )

	//fmt.Println("srcpix len: ",len(srcpix), srcdiff)
//	fmt.Println("dstpix len: ", len(dstpix), dstdiff)
	
	srcaddr := 0
	dstaddr := 0
	
	xspace := 0x10000 * srcwidth/dstwidth // must be > 1
	xrecip :=  int(int64(0x100000000) / int64(xspace))
	for y := 0; y < height; y++ {
		var accumulate [4]uint16
		xcounter := xspace
		for x:=0;x<srcwidth; x++ {
			if xcounter > 0x10000 {

				pad:= 1
				accumulate[0] += uint16(srcpix[srcaddr])
				srcaddr += pad
				accumulate[1] += uint16(srcpix[srcaddr])
				srcaddr += pad
				accumulate[2] += uint16(srcpix[srcaddr])
				srcaddr += pad
				accumulate[3] += uint16(srcpix[srcaddr])
				srcaddr += pad
				
				xcounter -= 0x10000
			} else {
				xfrac := 0x10000 - xcounter
				dstpixvalue := uint8( ((int(accumulate[0]) + (int(srcpix[srcaddr]) * xcounter ) >> 16) * xrecip ) >> 16 )
				dstpix[dstaddr] = byte( dstpixvalue & 0xff)
				dstaddr+=1
				
				dstpixvalue =  uint8( ((int(accumulate[1]) + (int(srcpix[srcaddr+1])*xcounter ) >> 16 ) * xrecip ) >> 16)
				dstpix[dstaddr] = byte( dstpixvalue & 0xff )
				dstaddr+=1

				dstpixvalue = uint8( (( int(accumulate[2]) + (int(srcpix[srcaddr+2]) *xcounter ) >> 16 ) * xrecip ) >> 16)
				dstpix[dstaddr] = byte(dstpixvalue & 0xff)
				dstaddr+=1

				dstpixvalue = uint8( (( int(accumulate[3]) + (int(srcpix[srcaddr+3]) *xcounter ) >> 16 ) * xrecip ) >> 16)
				dstpix[dstaddr] = byte(dstpixvalue & 0xff)
				dstaddr+=1

				accumulate[0] = uint16( (int(srcpix[srcaddr]) * xfrac) >> 16)
				srcaddr +=1
				accumulate[1] = uint16( (int(srcpix[srcaddr]) * xfrac) >> 16)
				srcaddr +=1
				accumulate[2] = uint16( (int(srcpix[srcaddr]) * xfrac) >> 16)
				srcaddr +=1
				accumulate[3] = uint16( (int(srcpix[srcaddr]) * xfrac) >> 16)
				srcaddr +=1					
 				xcounter = xspace - xfrac
			}
			
		}
		srcaddr += srcdiff
		dstaddr += dstdiff
	}
	
}

func filter_shrink_Y_ONLYC(srcpix []byte, dstpix []byte,  width,  srcpitch,  dstpitch, srcheight, dstheight int) {
		
	var templine []uint16
	templine = make([]uint16, dstpitch*2)
	srcdiff := srcpitch - (width * 4)
	dstdiff := dstpitch - (width * 4)
	/*
	fmt.Println("srcpix len: ",len(srcpix), srcdiff)
	fmt.Println("dstpix len: ", len(dstpix), dstdiff,dstpitch)
	*/
	
	yspace := 0x10000 * srcheight/dstheight // must be > 1
	yrecip := int(int64(0x100000000) / int64(yspace)) // int may overflow
	ycounter := yspace

	templine_addr := 0

	srcaddr := 0
	dstaddr := 0
	
	for y := 0; y < srcheight; y++ {
		templine_addr = 0
		if ycounter > 0x10000 {
			for x:=0;x<width;x++ {
				pad:=1
				templine[templine_addr] += uint16(srcpix[srcaddr])
				srcaddr+=pad
				templine_addr+=pad
				templine[templine_addr] += uint16(srcpix[srcaddr])
				srcaddr+=pad
				templine_addr+=pad
				templine[templine_addr] += uint16(srcpix[srcaddr])
				srcaddr+=pad
				templine_addr+=pad
				templine[templine_addr] += uint16(srcpix[srcaddr])
				srcaddr+=pad
				templine_addr+=pad
			}
			ycounter -= 0x10000
		}else {
			yfrac := 0x10000 - ycounter
			for x:=0; x< width; x++ {
				
				dstpixvalue := uint8( (( int(templine[templine_addr]) + (int(srcpix[srcaddr]) *ycounter ) >> 16 ) * yrecip ) >> 16)
				dstpix[dstaddr] = byte(dstpixvalue & 0xff)
				dstaddr+=1
				srcaddr+=1
				templine_addr +=1

				dstpixvalue = uint8( (( int(templine[templine_addr]) + (int(srcpix[srcaddr]) *ycounter ) >> 16 ) * yrecip ) >> 16)
				dstpix[dstaddr] = byte(dstpixvalue & 0xff)
				dstaddr+=1
				srcaddr+=1
				templine_addr +=1

				dstpixvalue = uint8( (( int(templine[templine_addr]) + (int(srcpix[srcaddr]) *ycounter ) >> 16 ) * yrecip ) >> 16)
				dstpix[dstaddr] = byte( dstpixvalue & 0xff)
				dstaddr+=1
				srcaddr+=1
				templine_addr +=1
				
				dstpixvalue = uint8( (( int(templine[templine_addr]) + (int(srcpix[srcaddr]) *ycounter ) >> 16 ) * yrecip ) >> 16)
				dstpix[dstaddr] = byte( dstpixvalue & 0xff)
				dstaddr+=1
				srcaddr+=1
				templine_addr +=1				
				
			}
			dstaddr += dstdiff
			templine_addr=0
			srcaddr -= 4 * width

			for x :=0; x<width;x++ {
				templine[templine_addr] = uint16( (int(srcpix[srcaddr]) * yfrac) >> 16)
				templine_addr+=1
				srcaddr += 1

				templine[templine_addr] = uint16( (int(srcpix[srcaddr]) * yfrac) >> 16)
				templine_addr+=1
				srcaddr += 1

				templine[templine_addr] = uint16( (int(srcpix[srcaddr]) * yfrac) >> 16)
				templine_addr+=1
				srcaddr += 1

				templine[templine_addr] = uint16( (int(srcpix[srcaddr]) * yfrac) >> 16)
				templine_addr+=1
				srcaddr += 1				
			}
			ycounter = yspace - yfrac
		}
		srcaddr += srcdiff
	}
}

/* this function implements a bilinear filter in the X-dimension */
func filter_expand_X_ONLYC(srcpix []byte, dstpix []byte, height, srcpitch, dstpitch, srcwidth,dstwidth int ) {
	dstdiff := dstpitch - (dstwidth * 4 )
	factorwidth := 4

	var xidx0  []int
	xidx0 = make([]int, dstwidth*4)
	var xmult0 []int
	xmult0 = make([]int, dstwidth*factorwidth)
	var xmult1 []int
	xmult1 = make([]int, dstwidth*factorwidth)
	/* Create multiplier factors and starting indices and put them in arrays */
	for x:=0;x<dstwidth;x++ {
		xidx0[x] = x*( srcwidth-1)/dstwidth
		xmult1[x] = 0x10000 * ((x * (srcwidth - 1)) % dstwidth) / dstwidth
		xmult0[x] = 0x10000 * xmult1[x]
	}

	srcaddr:=0
	dstaddr:=0
	/* Do the scaling in raster order so we don't trash the cache */
	for y:=0;y<height;y++ {
		srcrow0 := srcaddr + y*srcpitch
		for x:=0;x<dstwidth;x++ {
			src := srcrow0 + xidx0[x]*4
			xm0 := xmult0[x]
			xm1 := xmult1[x]
			dstvalue := uint8( ( int(srcpix[src+0])*xm0 + int(srcpix[src+4])*xm1 ) >> 16)
			dstpix[dstaddr] = byte(dstvalue & 0xff)
			dstaddr+=1

			dstvalue = uint8( ( int(srcpix[src+1])*xm0 + int(srcpix[src+5])*xm1 ) >> 16)
			dstpix[dstaddr] = byte(dstvalue & 0xff)
			dstaddr+=1
			
			dstvalue = uint8( ( int(srcpix[src+2])*xm0 + int(srcpix[src+6])*xm1 ) >> 16)
			dstpix[dstaddr] = byte(dstvalue & 0xff)
			dstaddr+=1

			dstvalue = uint8( ( int(srcpix[src+3])*xm0 + int(srcpix[src+7])*xm1 ) >> 16)
			dstpix[dstaddr] = byte(dstvalue & 0xff)
			dstaddr+=1			
			
		}
		dstaddr += dstdiff
	}
}


func filter_expand_Y_ONLYC(srcpix []byte, dstpix []byte, width, srcpitch, dstpitch, srcheight,dstheight int) {

	srcaddr:=0
	dstaddr:=0
	for y:=0; y<dstheight; y++ {
		yidx0 := y*(srcheight-1)/dstheight
		srcrow0 := srcaddr+yidx0 *srcpitch
		srcrow1 := srcrow0 + srcpitch

		ymult1 := 0x10000 * ((y * (srcheight - 1)) % dstheight) / dstheight
		ymult0 := 0x10000 - ymult1

		for x:=0; x<width; x++ {
			dstvalue := uint8(((int(srcpix[srcrow0]) * ymult0) + (int(srcpix[srcrow1]) * ymult1)) >> 16)
			dstpix[dstaddr] = byte( dstvalue & 0xff)
			srcrow0+=1
			srcrow1+=1
			dstaddr+=1

			dstvalue = uint8(((int(srcpix[srcrow0]) * ymult0) + (int(srcpix[srcrow1]) * ymult1)) >> 16)
			dstpix[dstaddr] = byte( dstvalue & 0xff)
			srcrow0+=1
			srcrow1+=1
			dstaddr+=1

			dstvalue = uint8(((int(srcpix[srcrow0]) * ymult0) + (int(srcpix[srcrow1]) * ymult1)) >> 16)
			dstpix[dstaddr] = byte( dstvalue & 0xff)
			srcrow0+=1
			srcrow1+=1
			dstaddr+=1

			dstvalue = uint8(((int(srcpix[srcrow0]) * ymult0) + (int(srcpix[srcrow1]) * ymult1)) >> 16)
			dstpix[dstaddr] = byte( dstvalue & 0xff)
			srcrow0+=1
			srcrow1+=1
			dstaddr+=1			
		}
	}
}

func convert_24_32(srcpix []byte, srcpitch int ,dstpix []byte, dstpitch ,width,height int ){
	srcdiff := srcpitch - (width * 3) // assume src bpp is 3
	dstdiff := dstpitch - (width * 4) //

	srcaddr:=0
	dstaddr:=0
	for y:=0; y<height; y++ {
		for x:=0; x<width; x++ {
			dstpix[dstaddr] = srcpix[srcaddr]
			dstaddr +=1
			srcaddr +=1

			dstpix[dstaddr] = srcpix[srcaddr]
			dstaddr +=1
			srcaddr +=1
			
			dstpix[dstaddr] = srcpix[srcaddr]
			dstaddr +=1
			srcaddr +=1

			dstpix[dstaddr] = 0xff
			dstaddr +=1
		}
		srcaddr += srcdiff
		dstaddr += dstdiff
	}
}

func convert_32_24(srcpix []byte, srcpitch int, dstpix []byte, dstpitch ,width,height int) {
	srcdiff := srcpitch - (width * 4)
	dstdiff := dstpitch - (width * 3)

	srcaddr:=0
	dstaddr:=0
	for y:=0; y<height; y++ {
		for x:=0; x<width;x++ {
			dstpix[dstaddr] = srcpix[srcaddr]
			dstaddr +=1
			srcaddr +=1

			dstpix[dstaddr] = srcpix[srcaddr]
			dstaddr +=1
			srcaddr +=1
			
			dstpix[dstaddr] = srcpix[srcaddr]
			dstaddr +=1
			srcaddr +=1

			srcaddr +=1
		}

		srcaddr += srcdiff
		dstaddr += dstdiff
	}
}

//pygame.transform.smoothscale(Surface, (width, height), DestSurface = None) -> Surface
// pygame transform.c because ARM has no MMX or SSE ,use filter_shrink_X_ONLYC GENEIC style to perform smoot scale
// pygame transform.c scalesmooth() L:1260

func scalesmooth(src *sdl.Surface, dst *sdl.Surface) {
	filter_shrink_X := filter_shrink_X_ONLYC
	filter_expand_X := filter_expand_X_ONLYC

	filter_shrink_Y := filter_shrink_Y_ONLYC
	filter_expand_Y := filter_expand_Y_ONLYC
	
	
	srcpix:= src.Pixels()
	dstpix:= dst.Pixels()

	var dst32 []byte
	var temppix []byte
	
	srcpitch := src.Pitch
	dstpitch := dst.Pitch

	srcwidth  := src.W
	srcheight := src.H
	dstwidth  := dst.W
	dstheight := dst.H

	bpp := src.BytesPerPixel()

	tempwidth := 0
	temppitch := 0
	tempheight := 0
	
	/* convert to 32-bit if necessary */
	if bpp == 3 {
		newpitch := int(srcwidth * 4)
		var newsrc []byte
		newsrc = make([]byte,newpitch * int(srcheight))
		convert_24_32(srcpix, int(srcpitch), newsrc, newpitch,int(srcwidth),int(srcheight))
		srcpix = newsrc
		srcpitch = int32(newpitch)
		/* create a destination buffer for the 32-bit result */
		dstpitch = dstwidth << 2 // << 2 equal *4
		dst32 = make([]byte, dstpitch*dstheight)
		dstpix = dst32
	}

	if srcwidth != dstwidth && srcheight != dstheight {
		tempwidth = int(dstwidth)
		temppitch = tempwidth << 2
		tempheight = int(srcheight)
		temppix = make([]byte,temppitch*tempheight)
	}

	/* Start the filter by doing X-scaling */
	if dstwidth < srcwidth { // shrink 
		if srcheight != dstheight {
			filter_shrink_X(srcpix, temppix, int(srcheight),int(srcpitch), temppitch, int(srcwidth),int(dstwidth))
		}else {
			filter_shrink_X(srcpix, dstpix,  int(srcheight),int(srcpitch), int(dstpitch),  int(srcwidth),int(dstwidth))
		}
	}else if dstwidth > srcwidth { // expand
		if srcheight != dstheight {
			filter_expand_X(srcpix, temppix, int(srcheight),int(srcpitch), temppitch, int(srcwidth),int(dstwidth))
		}else {
			filter_expand_X(srcpix, dstpix,  int(srcheight),int(srcpitch), int(dstpitch),  int(srcwidth),int(dstwidth))
		}	
	}

	/* Now do the Y scale */
	if dstheight < srcheight {
		if srcwidth != dstwidth {
			filter_shrink_Y(temppix, dstpix, tempwidth, temppitch, int(dstpitch),int( srcheight), int(dstheight))
		}else{
			filter_shrink_Y(srcpix, dstpix, int(srcwidth), int(srcpitch), int(dstpitch), int(srcheight), int(dstheight))
		}
	}else if dstheight > srcheight {
		if srcwidth != dstwidth {
			filter_expand_Y(temppix, dstpix, tempwidth, temppitch, int(dstpitch), int(srcheight), int(dstheight))
		}else {
			filter_expand_Y(srcpix, dstpix, int(srcwidth), int(srcpitch),int(dstpitch), int(srcheight), int(dstheight))
		}
	}

	if bpp == 3 {
		convert_32_24(dst32, int(dstpitch), dst.Pixels(),int(dst.Pitch), int(dstwidth), int(dstheight))
	}
	
}


//Public 
func SmoothScale(src_surf *sdl.Surface, new_width, new_height int ) *sdl.Surface {

	if new_width < 0 || new_height < 0 {
		panic("Cannot scale to negative size")
	}
	
	bpp := src_surf.BytesPerPixel()
	if bpp < 3 || bpp > 4 {
		panic("Only 24-bit or 32-bit surfaces can be smoothly scaled")
	}

	newsurf,_ := newsurf_fromsurf(src_surf,new_width,new_height)
	
	if int(newsurf.W) != new_width || int(newsurf.H) != new_height {
		panic("Destination surface not the given width or height.")
	}

	if src_surf.BytesPerPixel() != newsurf.BytesPerPixel() {
		panic("Source and destination surfaces need the same format.")
	}

	if (( new_width * bpp + 3) >> 2) > int(newsurf.Pitch) {
		panic("SDL Error: destination surface pitch not 4-byte aligned.")
	}

	if new_width > 0 && new_height > 0 {
		newsurf.Lock()
		src_surf.Lock()

		srcpix := src_surf.Pixels()
		dstpix := newsurf.Pixels()
		/* handle trivial case */
		if int(src_surf.W) == new_width && int(src_surf.H) == new_height {
			srcaddr:=0
			dstaddr:=0
			for y:=0;y<new_height;y++ {
				srcaddr = y*int(newsurf.Pitch)
				dstaddr = y*int(src_surf.Pitch)
				
				for x:=0;x<new_width*bpp;x++ {
					dstpix[dstaddr+x] = srcpix[srcaddr+x]
				}
			}
			
		}else {
			scalesmooth(src_surf,newsurf)
		}
		
		src_surf.Unlock()
		newsurf.Unlock()
	}
	
	return newsurf
}
