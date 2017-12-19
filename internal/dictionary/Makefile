PACKAGES := .

help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

test:      ## Live reload tests
	@reflex -g '*.go' go test -v

cov:       ## Runs code coverage
	go test -a -coverprofile=coverage.cov .
	go tool cover -html=coverage.cov
