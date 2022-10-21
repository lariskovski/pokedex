GO=/usr/local/go/bin/go
LDFLAGS=-ldflags="-s -w"

build:
	@GOARCH=amd64 GOOS=linux ${GO} build ${LDFLAGS} -o bin/get get/main.go
	@GOARCH=amd64 GOOS=linux ${GO} build ${LDFLAGS} -o bin/post post/main.go
	@GOARCH=amd64 GOOS=linux ${GO} build ${LDFLAGS} -o bin/put put/main.go
	@GOARCH=amd64 GOOS=linux ${GO} build ${LDFLAGS} -o bin/delete delete/main.go

shrink:
	@/usr/bin/upx --brute bin/*

create:
	serverless create -t aws-go-dep -p .

deploy: build
	serverless deploy --verbose

remove:
	serverless remove --verbose

clean:
	@${GO} clean
	@/usr/bin/rm -f bin/*