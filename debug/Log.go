package debug

import (
	"fmt"
)

type Source string

const System Source = "sys"
const Canvas Source = "canvas"
const GC Source = "gc"
const Loop Source = "loop"

func Log(msg string, src Source) {
	allow := src == "" ||
		src == "fps" ||
		//src == FPS ||
		//src == GC ||
		//src == Loop ||
		//src == System ||
		src == Canvas

	if !allow {
		return
	}

	fmt.Println("(" + string(src) + ")  " + msg)
}