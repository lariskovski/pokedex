GO=/usr/local/go/bin/go
LDFLAGS=-ldflags="-s -w"

METHOD=get

SRC_DIR=$(METHOD)
OBJ_DIR=obj
SRC_FILES=$(wildcard $(SRC_DIR)/*)
OBJ_FILES:=$(patsubst $(SRC_DIR)/%.go,$(OBJ_DIR)/%,$(SRC_FILES))

.PHONY: ${SRC_FILES}

all: $(OBJ_FILES)

$(OBJ_FILES): $(SRC_FILES)
	@GOARCH=amd64 GOOS=linux ${GO} build ${LDFLAGS} -o $@ $<
	@/usr/bin/upx --brute $@

zip:
	@find obj -maxdepth 1 -type f -execdir zip '{}.zip' '{}' \;

clean:
	@${GO} clean
	@/usr/bin/rm -f ${OBJ_DIR}/*