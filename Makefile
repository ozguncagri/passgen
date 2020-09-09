run:
	go run *.go

build:
	go build

buildCurrentProduction:
	go build -ldflags="-s -w"

buildDarwinProduction:
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o passgen
	zip passgen_darwin_amd64.zip passgen
	rm passgen

buildLinuxProduction:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o passgen
	zip passgen_linux_amd64.zip passgen
	rm passgen

buildWindowsProduction:
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o passgen.exe
	zip passgen_windows_amd64.zip passgen.exe
	rm passgen.exe

buildAllProduction: buildDarwinProduction buildLinuxProduction buildWindowsProduction