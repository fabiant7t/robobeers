package main

import (
	"machine"
	"time"
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.Low()
		time.Sleep(1 * time.Second)
		led.High()
		time.Sleep(1 * time.Second)
	}
}
