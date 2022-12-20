version="1.0.0"
echo 'building for windows...'
GOOS=windows GOARCH=amd64 go build -o bin/acorn-windows-$version.exe .