package debug

import (
	"strconv"
	"time"
)

type FPS struct {
	start_time int64
	frame_count int64
}

func NewFPS() *FPS {
	return &FPS{start_time: time.Now().Unix()}
}

// todo HAL has an fps limiter built in
func (fps *FPS) Tick() {
	fps.frame_count++

	if fps.frame_count % 60 != 0 {
		return
	}

	now := time.Now().Unix()
	framesPerSecond := int(fps.frame_count / (now - fps.start_time))

	Log("     FPS " + strconv.Itoa(framesPerSecond), "fps")
}



