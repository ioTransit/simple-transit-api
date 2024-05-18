.PHONY: tailwind-watch
tailwind-watch:
	./tailwindcss -i ./styles/custom.css -o ./styles/tailwind.css --watch

.PHONY: tailwind-build
tailwind-build:
	./tailwindcss -i ./styles/custom.css -o ./styles/tailwind.min.css --minify

.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: templ-watch
templ-watch:
	templ generate --watch
	
.PHONY: dev
dev:
	go build -o ./tmp ./main.go && air

.PHONY: build
build:
	make tailwind-build
	make templ-generate
	go build -ldflags "-X main.Environment=production" -o ./bin ./main.go

.PHONY: vet
vet:
	go vet ./...

.PHONY: staticcheck
staticcheck:
	staticcheck ./...

.PHONY: test
test:
	  go test -race -v -timeout 30s ./...
