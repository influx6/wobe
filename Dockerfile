FROM golang:alpine

MAINTAINER EWETUMO ALEXANDER 	<trinoxf@gmail.com>

COPY ./.docker_build/wobe .

CMD ["/wobe/wobe"]
