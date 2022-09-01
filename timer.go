package main

type Timer struct {
	DurationInSeconds float32
	ElapsedSeconds    float32
	Started           bool
}

func (timer *Timer) Reset() {
	timer.Started = false
	timer.ElapsedSeconds = 0
}

func (timer *Timer) Tick(frameTime float32) {
	if timer.Started {
		timer.ElapsedSeconds += frameTime
	}
}

func (timer *Timer) Expired() bool {
	return timer.Started && timer.ElapsedSeconds > timer.DurationInSeconds
}
