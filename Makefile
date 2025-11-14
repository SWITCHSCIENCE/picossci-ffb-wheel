TARGET=pico
TINYGO=tinygo
NAME=diy-ffb-wheel
TAGS=
-include .env
export TARGET
export TAGS

.PHONY: build all flash wait mon

build:
	mkdir -p build
	$(TINYGO) build -scheduler tasks -tags '$(TAGS)' -target $(TARGET) -o build/$(NAME).uf2 .

build-docker:
	mkdir -p build
	docker build --rm -t localhost/tinygo30 .
	docker run --rm -it -v $$(pwd -W):/app -w /app localhost/tinygo30

all: flash wait monitor

flash:
	$(TINYGO) flash -tags '$(TAGS)' -target $(TARGET)  .

wait:
	sleep 2

mon:
	$(TINYGO) monitor -tags '$(TAGS)' -target $(TARGET) 

gdb:
	$(TINYGO) gdb -x -tags '$(TAGS)' -target $(TARGET) -programmer=jlink 

server:
	"C:\Program Files\SEGGER\JLink\JLinkGDBServer.exe" -if swd -port 3333 -speed 4000 -device rp2040_m0_0 &

docker:
	docker compose up build 