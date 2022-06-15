package main

import (
	img "another-tiny-render/internal/image"

	"math/rand"
	"time"

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

func getRandColor() (uint8, uint8, uint8) {
	return uint8(rand.Intn(255)),
		uint8(rand.Intn(255)),
		uint8(rand.Intn(255))
}

func main() {

	// just for random colors, may remove later
	rand.Seed(time.Now().UnixNano())

	sdlGlobalState := initSdlGlobalState(512, 512, "sdl")
	defer destroySdlGlobalState(&sdlGlobalState)

	cnvs, _ := img.BuildCanvasFromFile("../../assets/512x512dude.jpg")
	// {
	// var (
	// r, g, b uint8
	// )
	// cnvs.DrawChecker(32)
	//
	// r, g, b = getRandColor()
	// cnvs.SetPenColor(r, g, b)
	// cnvs.BrasenhamLine(10, 10, 500, 402)
	//
	// r, g, b = getRandColor()
	// cnvs.SetPenColor(r, g, b)
	// cnvs.BrasenhamLine(400, 20, 40, 350)
	//
	// for i := 1; i < 7; i++ {
	// r, g, b = getRandColor()
	// cnvs.SetPenColor(r, g, b)
	// cnvs.BrasenhamCircle(256, 256, i*40)
	// }
	//
	// r, g, b = getRandColor()
	// cnvs.SetPenColor(r, g, b)
	// cnvs.DDALine(440, 110, 40, 426)
	//
	// }

	blured, _ := img.BuildEmptyCanvas(512, 512, 3)
	img.FastGaussianBlurRGB(cnvs.GetData(), blured.GetData(), 512, 512, 3, 8)

	texture, _ := sdlGlobalState.render.CreateTexture(sdl.PIXELFORMAT_RGB24,
		sdl.TEXTUREACCESS_TARGET, int32(blured.GetWidth()), int32(blured.GetHeight()))
	defer texture.Destroy()
	rect := sdl.Rect{X: 0, Y: 0, W: int32(blured.GetWidth()), H: int32(blured.GetHeight())}
	texture.Update(&rect, blured.GetData(), blured.GetWidth()*3)

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
