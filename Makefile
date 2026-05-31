.PHONY: build test lint sync-assets install clean

BIN := bin/propc

build: sync-assets
	@mkdir -p bin
	go build -o $(BIN) ./cmd/propc

test: sync-assets
	go test ./...

lint:
	go vet ./...
	gofmt -l . | tee /dev/stderr | (! read)

# Copy the canonical tikz files from material/ into src/assets/ so go:embed
# picks them up. main.tex reads from material/; the Go binary reads from
# src/assets/. Keep them in sync via this rule.
sync-assets:
	cp material/generator.tikz       src/assets/generator.tikz
	cp material/generator.tikzstyles src/assets/generator.tikzstyles

install: sync-assets
	go install ./cmd/propc

clean:
	rm -rf bin
