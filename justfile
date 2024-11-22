# web
# =================

build-web:
	cd web && just build
	mkdir -p dist
	cp -r web/dist/* dist/

# api
# =================

generate-api-docs: validate-api-spec
	mkdir -p dist/reference
	redocly build-docs --output=dist/reference/index.html api.yml

validate-api-spec:
	redocly lint api.yml

# fontdl
# =================

# TODO: Generate the client with go generate
compile-fontdl:
	oapi-codegen -generate client,types -o internal/api/api.go api.yml
	go build ./cmd/fontdl

build-fontdl NAME:
	oapi-codegen -generate client,types -o internal/api/api.go api.yml
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/{{NAME}}-linux-amd64 ./cmd/fontdl
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/{{NAME}}-macos-amd64 ./cmd/fontdl
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o bin/{{NAME}}-macos-arm64 ./cmd/fontdl

# This action is run by CI when releasing fontdl
package-fontdl:
	just build-fontdl "$(git describe --exact-match --tags)"

# builder
# =================

compile-builder:
	go build ./cmd/builder

# Generate WOFF2 and CSS files
generate-font-files: compile-builder
	./builder --input-dir=fonts/ --output-dir=dist/
	# Generate a master css file containing all font css files
	cat dist/*.css > dist/api/v1/download/_.css

# global
# ================

clean-dist:
	rm -rf dist/*

serve:
	miniserve --index index.html dist/

test:
	go test -coverprofile=coverage.out ./...

fmt:
	gofumpt -w .

build: clean-dist generate-font-files generate-api-docs build-web
