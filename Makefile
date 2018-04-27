#-----------------------------------------------------------------------------
# Global Variables
#-----------------------------------------------------------------------------

DOCKER_USER ?= $(DOCKER_USER)
DOCKER_PASS ?= 

DOCKER_BUILD_ARGS := --build-arg HTTP_PROXY=$(http_proxy) --build-arg HTTPS_PROXY=$(https_proxy)



APP_VERSION := develop

#-----------------------------------------------------------------------------
# MAIN
#-----------------------------------------------------------------------------

.PHONY: default build test publish clean build_local
default: depend build test publish

depend: 
	go get -v -t -d ./...
test:
	go test -v ./...
build_local:
	go build 
	docker build  $(DOCKER_BUILD_ARGS) -t $(DOCKER_USER)/demo-istio:$(APP_VERSION) .
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
	docker build $(DOCKER_BUILD_ARGS) -t $(DOCKER_USER)/demo-istio:$(APP_VERSION)  .
publish: 
	docker login -u $(DOCKER_USER) -p $(DOCKER_PASS)
	docker push $(DOCKER_USER)/demo-istio:$(APP_VERSION)
clean:
	rm -rf demo-istio
lint:
	go get -u github.com/alecthomas/gometalinter
	gometalinter ./...



