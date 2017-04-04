FROM golang:alpine

MAINTAINER EWETUMO ALEXANDER 	<trinoxf@gmail.com>

# RUN apk add --update lxc bridge lxc-templates git libpcap-dev pkgconfig

RUN sh -c "mkdir -p /wobe"

WORKDIR /wobe

COPY .docker_build/wobe/wobe ./wobe

CMD ["/wobe/wobe"]
