BINARY=go-core-kit
VERSION=0.4.0
DATE=`date +%FT%T%z`
.PHONY: init

default:
	@echo ${BINARY}
	@echo ${VERSION}
	@echo ${DATE}

init:
	@go generate
	@echo "[ok] generate"

upgrade:
	@go-mod-upgrade