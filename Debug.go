package bl

import (
	"time"
	"fmt"
)

var g_fps_start_time = time.Now().Unix()
var g_fps_frame int64 = 0

func fps() {
	g_fps_frame++

	if g_fps_frame % 60 == 0 {
		var now = time.Now().Unix()
		fmt.Println("FPS ", g_fps_frame / (now - g_fps_start_time))
	}
}

func debug(msg string) {
	if false {
		fmt.Println(msg)
	}
}