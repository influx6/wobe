FROM golang:latest

COPY . /go/src/github.com/influx6/wobe

WORKDIR /go/src/github.com/influx6/wobe

# Run go get for vendors
RUN go get -u -v

# Install gometalinter
RUN go get -u -v github.com/alecthomas/gometalinter

# Install missing lint tools
RUN gometalinter --install

# Run go linters
RUN gometalinter --deadline 4m --errors ./

# Run go tests
RUN go test -v
