env GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o build/passgen
zip build/passgen_darwin_amd64.zip build/passgen
rm build/passgen

env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o build/passgen
zip build/passgen_linux_amd64.zip build/passgen
rm build/passgen

env GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o build/passgen.exe
zip build/passgen_windows_amd64.zip build/passgen.exe
rm build/passgen.exe