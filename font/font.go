package font

import (
	"fmt"
	"bytes"
  
	"github.com/veandco/go-sdl2/ttf"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/cuu/gogame/color"
	"github.com/cuu/gogame/surface"	
)
/*
const (
    HINTING_NORMAL = int(C.TTF_HINTING_NORMAL)
    HINTING_LIGHT  = int(C.TTF_HINTING_LIGHT)
    HINTING_MONO   = int(C.TTF_HINTING_MONO)
    HINTING_NONE   = int(C.TTF_HINTING_NONE)
)
const (
    STYLE_NORMAL        = 0
    STYLE_BOLD          = 0x01
    STYLE_ITALIC        = 0x02
    STYLE_UNDERLINE     = 0x04
    STYLE_STRIKETHROUGH = 0x08
)
*/

var font_defaultname = "freesansbold.ttf"  // /usr/share/fonts/truetype/freefont/FreeSansBold.ttf

func Init() {
	if ttf.WasInit() == false {
		err := ttf.Init()
		if err != nil {
			panic( fmt.Sprintf("font Init failed %s",ttf.GetError()))
		}
	}
}

func Quit() {
	
	if ttf.WasInit() == true {
		ttf.Quit()
	}
	
}


func SysFont(name string, size int, bold bool, italic bool )  *ttf.Font {
	if ttf.WasInit() == false {
		panic("Init Font by using font.Init() first")
	}

	if size < 1 {
		size = 1
	}
	return nil
}

func Font(filename string, size int) *ttf.Font {

	if ttf.WasInit() == false {
		panic("Init Font by using font.Init() first")
	}

	if size < 1 {
		size = 1
	}
	
	fnt,err := ttf.OpenFont(filename,size)
	if err != nil {
		panic(fmt.Sprintf( "Open font %s failed %s", filename,err))
	}

	return fnt
}

func GetBold(fnt *ttf.Font) bool {
	cur_style := fnt.GetStyle()

	if cur_style & ttf.STYLE_BOLD  != 0 {
		return true
	}else {
		return false
	}
}
	
func SetBold(fnt *ttf.Font, bold_or_not bool ) {
	cur_style := fnt.GetStyle()
	
	if bold_or_not == true {
		cur_style = cur_style | ttf.STYLE_BOLD
		fnt.SetStyle( cur_style)
	}else {
		cur_style = cur_style & (^ttf.STYLE_BOLD)
		fnt.SetStyle( cur_style )
	}
}

func Size(fnt *ttf.Font, text string) (int,int) {
	w := 0
	h := 0
	w,h,err := fnt.SizeUTF8(text)
	if err != nil {
		fmt.Println("Get size of %s failed", text)
		w = 0
		h = 0
	}

	return w,h
}

//Return the height in pixels for a line of text with the font. When rendering multiple lines of text this is the recommended amount of space between lines.
func LineSize(fnt *ttf.Font) int {
	return fnt.LineSkip()
}


//Only UTF8 
func Render(fnt *ttf.Font,text string,antialias bool,col *color.Color, background *color.Color) *sdl.Surface {
	/*
	if antialias == true {
		fnt.SetHinting(ttf.HINTING_MONO)
	}
  */
	var surf *sdl.Surface
  var err error
	
	just_return := 0
	if text == "" || bytes.Equal( []byte(text),[]byte{0} ) {
		just_return = 1
	}

	if just_return  != 0 {
		height := fnt.Height()
		surf = surface.Surface(1,height)
		if background != nil {
			c := sdl.MapRGB(surf.Format,uint8(background.R),uint8(background.G),uint8(background.B))
			surf.FillRect(nil, c)
		}else {
			surf.SetColorKey(true,0)
		}
    return surf
	}
	
	if antialias == true {
		if background != nil {
			surf,err = fnt.RenderUTF8Shaded(text, col.ToSDL(), background.ToSDL())
			if err != nil {
				panic(fmt.Sprintf("%s Render failed  %s,%v,%d", fnt.FaceFamilyName(), ttf.GetError(),[]byte(text),len(text)))
			}			
		}else {
			surf,err = fnt.RenderUTF8Blended(text, col.ToSDL())
			if err != nil {
				panic(fmt.Sprintf("%s Render failed  %s,%v,%d", fnt.FaceFamilyName(), ttf.GetError(),[]byte(text),len(text)))
			}
		}
	
	}else {
		surf,err = fnt.RenderUTF8Solid(text, col.ToSDL())
		if err != nil {
			panic(fmt.Sprintf("%s Render failed  %s,%v,%d", fnt.FaceFamilyName(), ttf.GetError(),[]byte(text),len(text)))
		}
	}

	if antialias == false && background != nil && just_return == 0 {
		/* turn off transparancy */
		surf.SetColorKey(false,0) // false to disable colokey transparancy
		if surf.Format.Palette != nil {
			if surf.Format.Palette.Ncolors > 0 {
				surf.Format.Palette.Colors.R = uint8(background.R)
				surf.Format.Palette.Colors.R = uint8(background.G)
				surf.Format.Palette.Colors.R = uint8(background.B)
			}
		}
	}

	return surf 
}

