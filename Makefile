.PHONY: all clean test bench build

# 编译传入参数: 应用名称,默认iot-data
NAME:=flog
# 编译传入参数: Golang版本
GO_VERSION:=$(shell go version)
# 编译传入参数: 编译时间
BUILD_TIME:=$(shell date +'%FT%T%z')
# tag的版本号，自动根据tag填入
TAG_VERSION:=$(shell git describe --tags --always)
# 编译传入参数: 编译使用分支
BRANCH:=$(shell git rev-parse --abbrev-ref HEAD)

PARAMETER="\
-X 'main.name=$(NAME)' \
-X 'main.goVersion=$(GO_VERSION)' \
-X 'main.buildTime=$(BUILD_TIME)' \
-X 'main.tagVersion=$(TAG_VERSION)' \
-X 'main.branch=$(BRANCH)'

GOPROXY:="https://goproxy.cn,direct"

.DEFAULT: all
all: test

clean:

test: prepare
	@echo test...
	@go test -v ./...

bench: prepare
	@echo benchmark...
	@go test -run=^$$ -bench=.

prepare:
	@go mod tidy
