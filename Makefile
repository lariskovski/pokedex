GO=/usr/local/go/bin/go
LDFLAGS=-ldflags="-s -w"

METHOD=get

SRC_DIR=$(METHOD)
OBJ_DIR=$(SRC_DIR)
ZIP_DIR=zip
SRC_FILES=$(wildcard $(SRC_DIR)/*)
OBJ_FILES:=$(patsubst $(SRC_DIR)/%.go,$(OBJ_DIR)/%,$(SRC_FILES))

.PHONY: ${SRC_FILES} zip

all: $(OBJ_FILES)

$(OBJ_FILES): $(SRC_FILES)
	@GOARCH=amd64 GOOS=linux ${GO} build ${LDFLAGS} -o $@ $<
	@/usr/bin/upx --brute $@
	@mkdir -p zip
	@cd $(METHOD) && zip ../zip/$(METHOD).zip main

clean:
	@${GO} clean
	@/usr/bin/rm -f ${ZIP_DIR}/*
	@ find . -type f -name "main" -exec rm -f {} \;