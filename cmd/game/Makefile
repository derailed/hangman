VERSION  := 1.2
NAME     := game
BUILD    := $(shell cat /tmp/build)

help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

build: stamp ## Build go executable
	@GOOS=linux go build -ldflags "-w -X main.Build=`cat /tmp/build`" -a -tags netgo -o execs/$(NAME) main.go

img: build   ## Build Docker image
	docker build --rm -t quay.io/derailed/$(NAME):$(VERSION) .

push: img    ## Push image to Quay
	docker push quay.io/derailed/$(NAME)

stamp:
	@git rev-parse --short HEAD > /tmp/build

