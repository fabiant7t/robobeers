package main

import (
	"machine"
	"time"
)

const (
	period = 20e6 // 20000 Microseconds
)

func sleep() {
	time.Sleep(500 * time.Millisecond)
}

func main() {
	pin := machine.GPIO16
	pwm := machine.PWM0

	err := pwm.Configure(machine.PWMConfig{
		Period: period,
	})
	if err != nil {
		panic(err)
	}
	ch, err := pwm.Channel(machine.Pin(pin))
	if err != nil {
		panic(err)
	}

	for {
		// Center
		value := uint64(pwm.Top()) * 1500 / (period / 1000)
		pwm.Set(ch, uint32(value))
		sleep()

		// 90 degrees left
		value = uint64(pwm.Top()) * 2000 / (period / 1000)
		pwm.Set(ch, uint32(value))
		sleep()

		// Center
		value = uint64(pwm.Top()) * 1500 / (period / 1000)
		pwm.Set(ch, uint32(value))
		sleep()

		// 90 degrees right
		value = uint64(pwm.Top()) * 1000 / (period / 1000)
		pwm.Set(ch, uint32(value))
		sleep()
	}
}
