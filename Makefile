GO=/usr/local/go/bin/go
GOFLAGS=-ldflags="-s -w"

BINARY_NAME=main
PATH=get

.PHONY: listpath compile run zip clean

listpath:
	@/bin/echo $(shell echo $$PATH)

get/main: get/main.go
	@GOARCH=amd64 GOOS=linux ${GO} build ${GOFLAGS} -o $@ $^
	@/usr/bin/upx --brute $@

run: compile
	@./${PATH}${BINARY_NAME}

get/main.zip: compile
	@/usr/bin/zip $@ ${PATH}/${BINARY_NAME}

clean:
	@${GO} clean
	@/usr/bin/rm -f ${PATH}/${BINARY_NAME}
	@/usr/bin/rm -f ${PATH}/${BINARY_NAME}.upx
	@/usr/bin/rm -f ${PATH}/${BINARY_NAME}.zip