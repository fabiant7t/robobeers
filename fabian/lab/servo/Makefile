GOTOOLCHAIN=go1.21.7+auto

build:
	docker run \
	  --rm \
	  -v $(shell pwd):/src \
	  -w /src \
	  tinygo/tinygo:0.30.0 tinygo build \
	  -o /src/main.hex \
	  -size=short \
	  -target=pico

flash:
	GOTOOLCHAIN=${GOTOOLCHAIN} tinygo flash \
	  -target pico \
	  main.go
