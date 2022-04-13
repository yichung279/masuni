GO=go
GOOS=linux
GOARCH=amd64

# TODO: deploy
# deploy/presign_url_handler:

build/presign_url_handler:
	mkdir -p build/backend/presign_url/
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build -o ./build/backend/presign_url/handler github.com/yichung279/masuni/backend/presign_url/handler

test/presign_url_handler:
	go test ./backend./presign_url/...

lint/presign_url_handler:
	cd backend && golangci-lint run ./presign_url/...

test/backend:
	go test ./backend/...

lint/backend:
	cd backend && golangci-lint ./...

clean:
	rm -rf build
