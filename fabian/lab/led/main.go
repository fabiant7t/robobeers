package main

import (
	"machine"
	"time"
)

func main() {
	outputConfig := machine.PinConfig{Mode: machine.PinOutput}
	redLED := machine.D13
	redLED.Configure(outputConfig)
	for {
		dur := time.Duration(250 * time.Millisecond)
		redLED.Low()
		time.Sleep(dur)
		redLED.High()
		time.Sleep(dur)
	}
}
