FROM debian:bookworm-slim

RUN apt update && apt install -y \
  curl git golang
RUN curl -LO https://github.com/tinygo-org/tinygo/releases/download/v0.30.0/tinygo_0.30.0_amd64.deb
RUN dpkg -i tinygo_0.30.0_amd64.deb
CMD [ "tinygo", "build", "-target", "pico", "-o", "build/diy-ffb-wheel.uf2", "." ]
