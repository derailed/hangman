PACKAGES := .

help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

test:      ## Live reload tests
	@reflex -g '*.go' go test -v

cover:
	@echo "mode: set" > coverage.cov ; \
	for package in $(PACKAGES) ; do \
		go test -a -coverprofile=tmp.cov ./$$package ; \
		cat tmp.cov | grep -v "mode: set" >> coverage.cov ; \
	done ; \
	rm tmp.cov ; \

cov: cover ## Runs code coverage
	@go tool cover -html=coverage.cov
