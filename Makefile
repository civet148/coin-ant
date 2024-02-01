#SHELL=/usr/bin/env bash

CLEAN:=
BINS:=
DATE_TIME=`date +'%Y%m%d %H:%M:%S'`
COMMIT_ID=`git rev-parse --short HEAD`

build:
	rm -f coin-ant
	go mod tidy && go build -ldflags "-s -w -X 'main.BuildTime=${DATE_TIME}' -X 'main.GitCommit=${COMMIT_ID}'" -o coin-ant cmd/main.go
.PHONY: build
BINS+=coin-ant

docker:
	docker build --build-arg GIT_USER=${GIT_USER} --build-arg GIT_PASSWORD=${GIT_PASSWORD} --tag coin-ant -f Dockerfile .
.PHONY: docker

# 检查环境变量
env-%:
	@ if [ "${${*}}" = "" ]; then \
	    echo "Environment variable $* not set"; \
	    exit 1; \
	fi

db2go:
	go install github.com/civet148/db2go@latest
.PHONY: db2go

clean:
	rm -rf $(CLEAN) $(BINS)
.PHONY: clean
