package main

import (
	"runtime"
	"fmt"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-hal-oob"
)

func initialize() {
}

func tick() {

	bl.Root()
	{
		bl.Pos(64,64)
		bl.Dim(800,600)
		bl.Color(.1,.1,.1)
		bl.FlagBorderAll()

		bl.Font("arial", 10)
		bl.FontColor(1,1,1)
		bl.FontNudge(3,3)
		bl.Label("Hello world")

		bl.BorderThickness(bl.FourTwosInt)
		bl.BorderColor(1,1,1)
	}
	bl.End()
}

func uninit() {
}

func init() {
	runtime.LockOSThread()
}

func main() {
	bl.Start( haloob.New(), 1280, 1024, "Bellina v0.2", initialize, tick, uninit )

	fmt.Println("bye!")
}


