version="1.0.0"
echo 'building for Windows...'
rm bin/*.exe
GOOS=windows GOARCH=amd64 go build -o bin/acorn-$version.exe .