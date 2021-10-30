PKG = github.com/k1LoW/gh-grep
COMMIT = $$(git describe --tags --always)
OSNAME=${shell uname -s}
ifeq ($(OSNAME),Darwin)
	DATE = $$(gdate --utc '+%Y-%m-%d_%H:%M:%S')
else
	DATE = $$(date --utc '+%Y-%m-%d_%H:%M:%S')
endif

export GO111MODULE=on
export CGO_ENABLED=0

BUILD_LDFLAGS = -X $(PKG).commit=$(COMMIT) -X $(PKG).date=$(DATE)

default: test

ci: depsdev test sec

test:
	go test ./... -coverprofile=coverage.out -covermode=count

sec:
	gosec ./...

lint:
	golangci-lint run ./...

build:
	go build -ldflags="$(BUILD_LDFLAGS)"

depsdev:
	go install github.com/Songmu/ghch/cmd/ghch@v0.10.2
	go install github.com/Songmu/gocredits/cmd/gocredits@v0.2.0
	go install github.com/securego/gosec/v2/cmd/gosec@v2.8.1

prerelease:
	git pull origin main --tag
	go mod tidy
	ghch -w -N ${VER}
	gocredits . > CREDITS
	git add CHANGELOG.md CREDITS go.mod go.sum
	git commit -m'Bump up version number'
	git tag ${VER}

release:
	git push origin main --tag
	goreleaser --rm-dist

.PHONY: default test
