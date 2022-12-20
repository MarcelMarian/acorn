version="1.0.0"
echo 'building for Windows...'
GOOS=windows GOARCH=amd64 go build -o bin/acorn-$version.exe .