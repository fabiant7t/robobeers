package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/servo"
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

	a, err := servo.NewArray(pwm)
	if err != nil {
		panic(err)
	}

	s, err := a.Add(pin)
	if err != nil {
		panic(err)
	}

	for {
		s.SetMicroseconds(1500)
		time.Sleep(1 * time.Second)
		s.SetMicroseconds(1600)
		time.Sleep(1 * time.Second)
		s.SetMicroseconds(1700)
		time.Sleep(1 * time.Second)
	}
}
