GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
ROOT_CMD=$(shell pwd)
DOCKER_BUILD=$(shell pwd)/.docker_build
DOCKER_CMD=$(DOCKER_BUILD)/wobe

$(DOCKER_CMD): clean
	mkdir -p $(DOCKER_BUILD)
	$(GO_BUILD_ENV) go build -v -o $(DOCKER_CMD) .

$(INSTALL_NOW_CMD):
	npm install -g now

clean:
	rm -rf $(DOCKER_BUILD)

build: $(DOCKER_CMD)

heroku:
	heroku create wobe
	heroku container:push web
	cd $(ROOT_CMD)

now:
	$(shell now)
	cd $(ROOT_CMD)
