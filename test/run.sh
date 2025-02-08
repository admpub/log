export GOOS=linux
export GOARCH=amd64
go build -o ./${GOOS}_${GOARCH}.test .