So my journey starts here. I've been into software for more than 30 years by now, came across a lot of languages and am currently using Go and Python most of the time (in that order).

So a natural fit seems to be [[TinyGo]] on a [[ESP32]] board. 

## Step 1: How to flash [[TinyGo]] on a ESP-WROOM_32?

I will use the[Blinking LED tutorial](https://tinygo.org/docs/tutorials/blinky/) as the basis in `fabian/lab/blinky`.

I am using arch btw.

```sh
sudo pacman -S tinygo
# tinygo v0.30.0-1: 1312 MiB download/2141 MiB install!
```

That's sufficiant to compile to [[WebAssembly]], but `avrdude` is required for microcontroller targets:

```sh
sudo pacman -S avrdude
```

TinyGo (v0.30.0) requires Go version 1.18 to 1.21, but latest release is 1.22 (and I'm using it). So I need to install the latest v1.21 release next the current one.

```sh
go install golang.org/dl/go1.21.7@latest
go1.21.7 download
[...]
Success. You may now run 'go1.21.7'

alias go=$(which go1.21.7)
```

Edit:
I can use tinygo when setting the go toolchain to v1.21.7, see
https://go.dev/doc/toolchain
```sh
go env -w GOTOOLCHAIN=go1.21.7+auto
```

TinyGo is compiled with the current go version, so it is not working this way. Trying docker to build the firmware:

```sh
docker run \
  --rm \
  -v $(pwd):/src \
  -w /src \
  tinygo/tinygo:0.30.0 tinygo build \
  -o /src/main.hex \
  -size=short \
  -target=esp32-coreboard-v2
```

That one errors out with `error: ROM segments are non-contiguous: /tmp/tinygo2299494242/main`. Using `feather-nrf52840` as the target works. So I'm searching my shelf for microcontrollers that I own and try to find one that does compile.

Working:
- [arduino-nano](https://tinygo.org/docs/reference/microcontrollers/arduino-nano/)
- [arduino-mega2560](https://tinygo.org/docs/reference/microcontrollers/arduino-mega2560/)

So I'll continue with the Mega2560:

```sh
docker run \
  --rm \
  -v $(pwd):/src \
  -w /src \
  tinygo/tinygo:0.30.0 tinygo build \
  -o /src/main.hex \
  -size=short \
  -target=arduino-mega2560
```

The following step might require [[Serial Access for Arduino]]

```sh
sudo avrdude \
  -p atmega2560 \
  -c wiring \
  -P /dev/ttyACM0 \
  -b 115200 \
  -U flash:w:main.hex
```

That one blinks!

I've added a Makefile, the steps are `make build` and `make flash`.

Unfortunately, [[TinyGo]] does not support [[PWM]] yet, which is required to control a [[Servo]]. I've ordered some [[Raspberry Pi Pico]] and a [[Seeed Xiao BLE]] and will have to wait a few days for PWM support.

## Step 2: LED with Resistor

Connecting a LED to D13 (with a resistor between 220 and 1000 Ohm)

## Step 3: LED with Button and Resistor

LED still on D13, Button on D2 and 10K resistor.
The resistor prevents signal pin from floating (not 1 or 0, but rapidly changing in between).

## Step 4: LED with Button with internal Resistor

LED still on D13, Button still on D2. Internal machine.PinInputPullup resistor.

## Step n: Add a SG90 PWM motor

This is basically the [Using PWM](https://tinygo.org/docs/tutorials/pwm/) tutorial.
[[PWM]] is the abbreviation for "[[Pulse-Width Modulation]]".
