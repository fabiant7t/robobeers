build:
	docker run \
	  --rm \
	  -v $(shell pwd):/src \
	  -w /src \
	  tinygo/tinygo:0.30.0 tinygo build \
	  -o /src/main.hex \
	  -size=short \
	  -target=arduino-mega2560

flash:
	sudo avrdude \
	  -p atmega2560 \
	  -c wiring \
	  -P /dev/ttyACM0 \
	  -b 115200 \
	  -U flash:w:main.hex
