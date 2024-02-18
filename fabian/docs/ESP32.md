[Core board](https://tinygo.org/docs/reference/microcontrollers/esp32-coreboard-v2/) with interface and pin docs

Flashing should be like that:
```sh
tinygo flash \
  -target=esp32-coreboard-v2 \
  -port=/dev/ttyUSB0 \
  examples/blinky1
```

