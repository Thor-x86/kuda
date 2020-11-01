all: clean win_x64 win_x32 linux_x64 linux_x32 linux_arm64 linux_arm32 mac

.PHONY: test demo

demo:
	rm -rf demo/dist
	cd demo && npm install && npm run build
	go build
	./kuda demo/dist

test:
	go test

clean:
	rm -f kuda
	rm -rf build/*
	rm -rf demo/dist

win_x64:
	GOOS=windows GOARCH=amd64 go build -o build/kuda-win_x64.exe

win_x32:
	GOOS=windows GOARCH=386 go build -o build/kuda-win_x32.exe

linux_x64:
	GOOS=linux GOARCH=amd64 go build -o build/kuda-linux_x64

linux_x32:
	GOOS=linux GOARCH=386 go build -o build/kuda-linux_x32

linux_arm64:
	GOOS=linux GOARCH=arm64 go build -o build/kuda-linux_arm64

linux_arm32:
	GOOS=linux GOARCH=arm go build -o build/kuda-linux_arm32

mac:
	GOOS=darwin GOARCH=amd64 go build -o build/kuda-mac