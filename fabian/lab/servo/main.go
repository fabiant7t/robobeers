package main

import (
	"machine"
	"time"
)

var (
	period uint64 = 2000000 // in nanoseconds, it's 10^9 / frequency (50 Hz) in Austria
)

func main() {
	// This program is specific to the Raspberry Pi Pico.
	pin := machine.D9
	pwm := machine.Timer1

	// Configure the PWM with the given period.
	pwm.Configure(machine.PWMConfig{
		Period: period,
	})

	ch, err := pwm.Channel(pin)
	if err != nil {
		println(err.Error())
		return
	}

	for {
		for i := 1; i < 255; i++ {
			// This performs a stylish fade-out blink
			pwm.Set(ch, pwm.Top()/uint32(i))
			time.Sleep(time.Millisecond * 5)
		}
	}
}
