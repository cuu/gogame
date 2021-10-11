package draw

import (
	"fmt"
	//	"math"
	"github.com/cuu/gogame/color"
	"github.com/veandco/go-sdl2/sdl"
)

func draw_ellipse(dst *sdl.Surface, x int, y int, rx int, ry int, color *color.Color) {

	if rx == 0 && ry == 0 {
		pixel(dst, color, x, y)
		return
	}

	if rx == 0 { /* Special case for rx=0 - draw a vline */
		drawvertlineclip(dst, color, x, y-ry, y+ry)
		return
	}

	if ry == 0 { /* Special case for ry=0 - draw a hline */
		drawhorzlineclip(dst, color, x-rx, y, x+rx)
		return
	}

	var oh, oi, oj, ok int
	oh = 0xffff
	oi = 0xffff
	oj = 0xffff
	ok = 0xffff

	var ix, iy int
	var h, i, j, k int
	var xmh, xph, ypk, ymk int
	var xmi, xpi, ymj, ypj int
	var xmj, xpj, ymi, ypi int
	var xmk, xpk, ymh, yph int

	if rx > ry {
		ix = 0
		iy = rx * 64
		for {
			h = (ix + 16) >> 6
			i = (iy + 16) >> 6
			j = (h * ry) / rx
			k = (i * ry) / rx

			if ((ok != k) && (oj != k)) || ((oj != j) && (ok != j)) || (k != j) {
				xph = x + h - 1
				xmh = x - h
				if k > 0 {
					ypk = y + k - 1
					ymk = y - k
					if h > 0 {
						pixel(dst, color, xmh, ypk)
						pixel(dst, color, xmh, ymk)
					}
					pixel(dst, color, xph, ypk)
					pixel(dst, color, xph, ymk)
				}
				ok = k
				xpi = x + i - 1
				xmi = x - i
				if j > 0 {
					ypj = y + j - 1
					ymj = y - j
					pixel(dst, color, xmi, ypj)
					pixel(dst, color, xpi, ypj)
					pixel(dst, color, xmi, ymj)
					pixel(dst, color, xpi, ymj)
				}
				oj = j
			}
			ix = ix + iy/rx
			iy = iy - ix/rx

			if i <= h {
				break
			}
		}
	} else {

		ix = 0
		iy = ry * 64
		for {
			h = (ix + 32) >> 6
			i = (iy + 32) >> 6
			j = (h * rx) / ry
			k = (i * rx) / ry

			if ((oi != i) && (oh != i)) || ((oh != h) && (oi != h) && (i != h)) {
				xmj = x - j
				xpj = x + j - 1
				if i > 0 {
					ypi = y + i - 1
					ymi = y - i
					if j > 0 {
						pixel(dst, color, xmj, ypi)
						pixel(dst, color, xmj, ymi)
					}
					pixel(dst, color, xpj, ypi)
					pixel(dst, color, xpj, ymi)
				}
				oi = i
				xmk = x - k
				xpk = x + k - 1
				if h > 0 {
					yph = y + h - 1
					ymh = y - h
					pixel(dst, color, xmk, yph)
					pixel(dst, color, xpk, yph)
					pixel(dst, color, xmk, ymh)
					pixel(dst, color, xpk, ymh)
				}
				oh = h
			}
			ix = ix + iy/ry
			iy = iy - ix/ry

			if i <= h {
				break
			}

		}
	}
}

func draw_fillellipse(dst *sdl.Surface, x int, y int, rx int, ry int, color *color.Color) {
	var ix, iy int
	var h, i, j, k int
	var oh, oi, oj, ok int

	if rx == 0 && ry == 0 {
		pixel(dst, color, x, y)
		return
	}

	if rx == 0 {
		drawvertlineclip(dst, color, x, y-ry, y+ry)
		return
	}

	if ry == 0 {
		drawhorzlineclip(dst, color, x-rx, y, x+rx)
		return
	}

	oh = 0xffff
	oh = 0xffff
	oj = 0xffff
	ok = 0xffff

	if rx >= ry {
		ix = 0
		iy = rx * 64

		for {
			h = (ix + 8) >> 6
			i = (iy + 8) >> 6
			j = (h * ry) / rx
			k = (i * ry) / rx
			if (ok != k) && (oj != k) && (k < ry) {
				drawhorzlineclip(dst, color, x-h, y-k-1, x+h-1)
				drawhorzlineclip(dst, color, x-h, y+k, x+h-1)
				ok = k
			}
			if (oj != j) && (ok != j) && (k != j) {
				drawhorzlineclip(dst, color, x-i, y+j, x+i-1)
				drawhorzlineclip(dst, color, x-i, y-j-1, x+i-1)
				oj = j
			}
			ix = ix + iy/rx
			iy = iy - ix/rx

			if i <= h {
				break
			}
		}
	} else {
		ix = 0
		iy = ry * 64

		for {
			h = (ix + 8) >> 6
			i = (iy + 8) >> 6
			j = (h * rx) / ry
			k = (i * rx) / ry

			if (oi != i) && (oh != i) && (i < ry) {
				drawhorzlineclip(dst, color, x-j, y+i, x+j-1)
				drawhorzlineclip(dst, color, x-j, y-i-1, x+j-1)
				oi = i
			}
			if (oh != h) && (oi != h) && (i != h) {
				drawhorzlineclip(dst, color, x-k, y+h, x+k-1)
				drawhorzlineclip(dst, color, x-k, y-h-1, x+k-1)
				oh = h
			}

			ix = ix + iy/ry
			iy = iy - ix/ry

			if i <= h {
				break
			}
		}
	}
}

func Circle(surf *sdl.Surface, color *color.Color, x, y, radius, border_width int) {

	if radius < 0 {
		fmt.Println("Circle negative radius")
		return
	}

	if border_width < 0 {
		fmt.Println("Circle negative border width")
		return
	}

	if border_width > radius {
		fmt.Println("Circle border width greater than radius")
		return
	}

	if border_width == 0 {
		draw_fillellipse(surf, x, y, radius, radius, color)
	} else {
		for loop := 0; loop < border_width; loop++ {
			draw_ellipse(surf, x, y, radius-loop, radius-loop, color)
			/* To avoid moiré pattern. Don't do an extra one on the outer ellipse.
			   We draw another ellipse offset by a pixel, over drawing the missed
			   spots in the filled circle caused by which pixels are filled.
			*/
			if border_width > 1 && loop > 0 {
				draw_ellipse(surf, x+1, y, radius-loop, radius-loop, color)
			}
		}
	}
}
