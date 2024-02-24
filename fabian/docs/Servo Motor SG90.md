[Datasheet](http://www.ee.ic.ac.uk/pcheung/teaching/DE1_EE/stores/sg90_datasheet.pdf)
[Servo driver](https://pkg.go.dev/tinygo.org/x/drivers/servo)

Uses 4.8 to 6 V

Ground is brown
VCC is red
PWM is orange

50 Hertz -> 20 ms (One PWM period)

Duty cycle:
0 degrees (center) 1.5ms
+90 degrees 2ms
-90 degrees 1ms

Wait for the rest of the period.