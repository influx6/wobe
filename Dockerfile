FROM golang:alpine

MAINTAINER EWETUMO ALEXANDER 	<trinoxf@gmail.com>

COPY ./.docker_build/wobe /bin

CMD ["/bin/wobe"]
