GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
ROOT_CMD=$(shell pwd)
DOCKER_BUILD=$(shell pwd)/.docker_build
DOCKER_CMD=$(DOCKER_BUILD)/wobe

build: clean
	mkdir -p $(DOCKER_BUILD)
	$(GO_BUILD_ENV) go build -v -o $(DOCKER_CMD) .
	docker build -t wobe ./

clean:
	rm -rf $(DOCKER_BUILD)

$(build)
