package bl

import (
	"time"
	"fmt"
	"strconv"
)

var g_fps_start_time = time.Now().Unix()
var g_fps_frame int64 = 0

func fps() {

	g_fps_frame++

	if g_fps_frame % 60 == 0 {
		var now = time.Now().Unix()
		debug("     FPS " + strconv.Itoa(int(g_fps_frame / (now - g_fps_start_time))), "fps")
	}
}

func debug(msg, src string) {

	log := src == "canvas" ||
		   //src == "fps" ||
		   src == "gc" ||
		   //src == "loop" ||
		   src == "sys" ||
		   src == ""

	//log = false

	if log {
		fmt.Println("(" + src + ")  " + msg)
	}
}