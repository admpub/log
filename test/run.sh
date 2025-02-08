export GOOS=linux
export GOARCH=amd64
tinygo build -o ./${GOOS}_${GOARCH}.test .