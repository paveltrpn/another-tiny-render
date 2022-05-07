package main

import (
	img "tiny-render-go/internal/image"

	"github.com/veandco/go-sdl2/sdl"
)

type sdlGlobalState_s struct {
	window *sdl.Window
	screen *sdl.Surface
	render *sdl.Renderer

	wnd_width  int
	wnd_height int
	name       string

	run bool
}

func initSdlGlobalState(width, height int, name string) sdlGlobalState_s {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		println("initSdlGlobalState(): ERROR while sdl init!")
		panic(err)
	}

	window, err := sdl.CreateWindow(name,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		int32(width), int32(height), sdl.WINDOW_SHOWN)
	if err != nil {
		println("initSdlGlobalState(): ERROR while window creation!")
		panic(err)
	}

	surface, err := window.GetSurface()
	if err != nil {
		println("initSdlGlobalState(): ERROR while screen sourface creation!")
		panic(err)
	}

	surface.FillRect(nil, 100)

	render, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_SOFTWARE)
	if err != nil {
		println("initSdlGlobalState(): ERROR while render creation!")
		panic(err)
	}

	return sdlGlobalState_s{window: window,
		screen:     surface,
		render:     render,
		wnd_width:  width,
		wnd_height: height,
		name:       name,
		run:        true}
}

func destroySdlGlobalState(state *sdlGlobalState_s) {
	state.window.Destroy()
	sdl.Quit()
}

func main() {
	sdlGlobalState := initSdlGlobalState(800, 600, "sdl")
	defer destroySdlGlobalState(&sdlGlobalState)

	cnvs, _ := img.BuildCanvas(512, 512, 3)
	cnvs.BrasenhamLine(10, 10, 500, 402)
	cnvs.BrasenhamLine(400, 20, 40, 350)

	texture, _ := sdlGlobalState.render.CreateTexture(sdl.PIXELFORMAT_RGB24,
		sdl.TEXTUREACCESS_TARGET, int32(cnvs.GetWidth()), int32(cnvs.GetHeight()))
	defer texture.Destroy()
	rect := sdl.Rect{X: 0, Y: 0, W: int32(cnvs.GetWidth()), H: int32(cnvs.GetHeight())}
	texture.Update(&rect, cnvs.GetDataPtr(), cnvs.GetWidth()*3)

	sdlGlobalState.window.UpdateSurface()

	for sdlGlobalState.run {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				sdlGlobalState.run = false

			case *sdl.KeyboardEvent:
				if t.Keysym.Sym == sdl.K_ESCAPE {
					sdlGlobalState.run = false
				}
			}
		}

		sdlGlobalState.render.Copy(texture, nil, &rect)
		sdlGlobalState.render.Present()
	}
}
