package gogame

import "github.com/veandco/go-sdl2/sdl"

const (
	
	SHOWN    = sdl.WINDOW_SHOWN
	OPENGL   = sdl.WINDOW_OPENGL
	HIDDEN   = sdl.WINDOW_HIDDEN 
    FIRSTHIDDEN = 0x20000000 // after SDL_WINDOW_VULKAN= 0x10000000 /**< window usable for Vulkan surface */
	FULLSCREEN = sdl.WINDOW_FULLSCREEN
	RESIZABLE  = sdl.WINDOW_RESIZABLE 
    ONTOP      = sdl.WINDOW_ALWAYS_ON_TOP //x11 only SDL >= 2.0.5
    INPUT_FOCUS = sdl.WINDOW_INPUT_FOCUS 
    
    UTILITY = sdl.WINDOW_UTILITY
    TOOLTIP = sdl.WINDOW_TOOLTIP
)
