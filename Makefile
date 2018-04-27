#-----------------------------------------------------------------------------
# Global Variables
#-----------------------------------------------------------------------------

DOCKER_USER:=
DOCKER_PASS:=

APP_VERSION := develop

#-----------------------------------------------------------------------------
# MAIN
#-----------------------------------------------------------------------------

.PHONY: default build 
default: depend build test publish

depend: 
	go get -v -t -d ./...
test:
	go test -v ./...
build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build 
	docker build -t $(DOCKER_USER)/demo-istio:$(APP_VERSION) .
build:
	go build
	docker build -t $(DOCKER_USER)/demo-istio:$(APP_VERSION)
publish: 
	docker login -u $(DOCKER_USER) -p $(DOCKER_PASS)
	docker push $(DOCKER_USER)/demo-istio:$(APP_VERSION)


