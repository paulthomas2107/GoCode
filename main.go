package main

import (
	"fmt"
	"os"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	if err := run() ; err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(2)
	}
}

func run() error  {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return fmt.Errorf("Could not start SDL: %v", err)
	}
	defer sdl.Quit()

	w, r, err := sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOWEVENT_SHOWN)

	if err != nil {
		return fmt.Errorf("Could not create window: %v", err)
	}
	
	defer w.Destroy()
	_= r

	return nil
}