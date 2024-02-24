[TinyGo reference](https://tinygo.org/docs/reference/microcontrollers/pico/)


Target: `pico`

Ordered 4 standard and 2 "wireless ones at Semaf Electronics in Vienna on 18th Feb 2024. Received them on 20th Feb 2024, yay!

Hardware (GPIO, UART, SPI, I2C, ADC, PWM, USBDevice) is fully supported by TinyGo. The wireless option is not.

The pico ships with a [UF2 Bootloader](https://github.com/Microsoft/uf2) preinstalled.

Flashing:

Keep "BOOTSEL" button pressed while plugging in the USB cable, then release it. A flash drive appears that you have to mount (also `lsusb` shows `Raspberry Pi Pico Boot` instead of `Raspberry Pi Pico`)

```sh
tinygo flash -target=pico main.go
```
## No more unplugging before flash
When frequent flashes are required, it's possible to short RUN  to GRND through a micro button, press it, press BOOTSEL button, release it and release BOOTSEL button. This way, you don't have to reattach the USB cable to boot to firmware flash mode.

## 5 Volts
The pico runs on 3.3 V, but there is the [VBUS (Pin 40)](https://www.raspberrypi.com/documentation/microcontrollers/images/pico-pinout.svg) where you can draw 5V when connected through micro USB. If you do not connect through USB, VSYS (pin 39) is used to provide the board with a 5V supply, which also means you cannot use VBUS anymore.