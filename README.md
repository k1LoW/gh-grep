# gh-grep

:octocat: Print lines matching a pattern in repositories using GitHub API

## Usage

``` console
$ gh grep func.*schema.Schema --include=**/*.go --owner k1LoW --repo tbls
k1LoW/tbls:cmd/doc.go:func withDot(s *schema.Schema, c *config.Config, force bool) (e error) {
k1LoW/tbls:cmd/doc.go:func outputErExists(s *schema.Schema, path string) bool {
k1LoW/tbls:config/config.go:func (c *Config) ModifySchema(s *schema.Schema) error {
k1LoW/tbls:config/config.go:func (c *Config) MergeAdditionalData(s *schema.Schema) error {
k1LoW/tbls:config/config.go:func (c *Config) FilterTables(s *schema.Schema) error {
k1LoW/tbls:config/config.go:func (c *Config) mergeDictFromSchema(s *schema.Schema) {
k1LoW/tbls:config/config.go:func excludeTableFromSchema(name string, s *schema.Schema) error {
[...]
```

( Do grep the codes (`**/*.go`) of [k1LoW/tbls](https://github.com/k1LoW/tbls) with the pattern `func.*schema.Schema` )

``` console
$ gh grep --help
Print lines matching a pattern in repositories using GitHub API

Usage:
  gh-grep [PATTERN] [flags]

Flags:
  -e, -- strings         match pattern
      --branch string    branch name
  -c, --count            show the number of matches instead of matching lines
      --exclude string   skip files and directories matching pattern
  -h, --help             help for gh-grep
  -i, --ignore-case      case insensitive matching
      --include string   search only files that match pattern (default "**/*")
  -n, --line-number      show line numbers
      --name-only        show only repository:filenames
  -o, --only-matching    show only matching parts of a line
      --owner string     repository owner or org
      --repo strings     repository name
      --repo-only        show only repositories
      --tag string       tag name
      --url              show URL
  -v, --version          version for gh-grep
```

#### :warning: Notice :warning:

**`gh-grep` is very slow because it does all its scanning through the GitHub API.**

**It is recommended to specify the `--include` option to get the results in a realistic time.**

## Examples

### List base Docker images used in the Dockerfile of the project root

``` console
$ gh grep ^FROM --include=Dockerfile --owner k1LoW
k1LoW/centve:Dockerfile:FROM centos:7
k1LoW/docker-alpine-pandoc-ja:Dockerfile:FROM frolvlad/alpine-glibc
k1LoW/docker-sshd:Dockerfile:FROM docker.io/alpine:3.9
k1LoW/gh-grep:Dockerfile:FROM debian:buster-slim
k1LoW/ghdag:Dockerfile:FROM debian:buster-slim
k1LoW/ghdag-action:Dockerfile:FROM ghcr.io/k1low/ghdag:v0.16.0
k1LoW/ghput:Dockerfile:FROM alpine:3.13
k1LoW/ghput-release-action:Dockerfile:FROM ghcr.io/k1low/ghput:v0.12.0
k1LoW/github-script-ruby:Dockerfile:FROM ghcr.io/k1low/github-script-ruby-base:v1.1.0
[...]
```

### List Actions you are using

``` console
$ gh grep uses: --include=.github/workflows/* --owner k1LoW | sed -e 's/.*uses:\s*//g' | sort | uniq -c
   9 ./
   1 EndBug/add-and-commit@v7
   2 actions/checkout@master
  10 actions/checkout@v1
  50 actions/checkout@v2
  18 actions/setup-go@v1
  21 actions/setup-go@v2
   4 aquasecurity/trivy-action@master
[...]
```

### Open the matched lines in a browser.

``` console
$ gh grep 'ioutil\.' --include=**/*.go --owner k1LoW --repo ghput --url
https://github.com/k1LoW/ghput/blob/main/gh/gh.go#L300
https://github.com/k1LoW/ghput/blob/main/gh/gh.go#L313
$ gh grep 'ioutil\.' --include=**/*.go --owner k1LoW --repo ghput --url | xargs open
```

*`open` command only works on macOS.

## Install

`gh-grep` can be installed as a standalone command or as [a GitHub CLI extension](https://cli.github.com/manual/gh_extension)

### Install as a GitHub CLI extension

``` console
$ gh extension install k1LoW/gh-grep
```

### Install as a standalone command

Run `gh-grep` instead of `gh grep`.

**deb:**

``` console
$ export GH_GREP_VERSION=X.X.X
$ curl -o gh-grep.deb -L https://github.com/k1LoW/gh-grep/releases/download/v$GH_GREP_VERSION/gh-grep_$GH_GREP_VERSION-1_amd64.deb
$ dpkg -i gh-grep.deb
```

**RPM:**

``` console
$ export GH_GREP_VERSION=X.X.X
$ yum install https://github.com/k1LoW/gh-grep/releases/download/v$GH_GREP_VERSION/gh-grep_$GH_GREP_VERSION-1_amd64.rpm
```

**apk:**

``` console
$ export GH_GREP_VERSION=X.X.X
$ curl -o gh-grep.apk -L https://github.com/k1LoW/gh-grep/releases/download/v$GH_GREP_VERSION/gh-grep_$GH_GREP_VERSION-1_amd64.apk
$ apk add gh-grep.apk
```

**homebrew tap:**

```console
$ brew install k1LoW/tap/gh-grep
```

**macports:**

```console
$ sudo port install gh-grep
```

Maintainer: @herbygillot

**manually:**

Download binary from [releases page](https://github.com/k1LoW/gh-grep/releases)

**go get:**

```console
$ go get github.com/k1LoW/gh-grep
```

**docker:**

```console
$ docker pull ghcr.io/k1low/gh-grep:latest
```
