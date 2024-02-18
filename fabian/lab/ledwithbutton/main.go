package main

import (
	"machine"
)

func main() {
	outputConfig := machine.PinConfig{Mode: machine.PinOutput}
	redLED := machine.D13
	redLED.Configure(outputConfig)

	inputConfig := machine.PinConfig{Mode: machine.PinInput}
	buttonInput := machine.D2
	buttonInput.Configure(inputConfig)

	for {
		if buttonInput.Get() {
			redLED.High()
			continue
		}
		redLED.Low()
	}
}
