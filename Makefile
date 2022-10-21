GO=/usr/local/go/bin/go
LDFLAGS=-ldflags="-s -w"

all:
	@GOARCH=amd64 GOOS=linux ${GO} build ${LDFLAGS} -o bin/get get/main.go
	@GOARCH=amd64 GOOS=linux ${GO} build ${LDFLAGS} -o bin/post post/main.go
	@GOARCH=amd64 GOOS=linux ${GO} build ${LDFLAGS} -o bin/put put/main.go
	@GOARCH=amd64 GOOS=linux ${GO} build ${LDFLAGS} -o bin/delete delete/main.go
	# @/usr/bin/upx --brute bin/*

create:
	serverless create -t aws-go-dep -p .

deploy:
	sls deploy --verbose

remove:
	sls remove --verbose

clean:
	@${GO} clean
	@/usr/bin/rm -f bin/*