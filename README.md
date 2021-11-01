# gh-grep

Print lines matching a pattern in repositories using GitHub API

## Usage

Do grep the codes (`**/*.go`) of [k1LoW/tbls](https://github.com/k1LoW/tbls) with the pattern `^import` .

``` console
$ gh grep ^import --include=**/*.go --owner k1LoW --repo tbls
```

## Example

### List Actions you are using

``` console
$ gh grep uses: --include=.github/workflows/* --owner [Your Login] | sed -e 's/.*uses:\s*//g' | sort | uniq -c
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

## Install

`gh-grep` can be installed as a standalone command or as [a GitHub CLI extension](https://cli.github.com/manual/gh_extension)

### Install as a GitHub CLI extension

``` console
$ gh extension install k1LoW/gh-grep
```

### Install as a standalone command

Run `gh-grep` instead of `gh grep`.

**deb:**

Use [dpkg-i-from-url](https://github.com/k1LoW/dpkg-i-from-url)

``` console
$ export GH-GREP_VERSION=X.X.X
$ curl -L https://git.io/dpkg-i-from-url | bash -s -- https://github.com/k1LoW/gh-grep/releases/download/v$GH-GREP_VERSION/gh-grep_$GH-GREP_VERSION-1_amd64.deb
```

**RPM:**

``` console
$ export GH-GREP_VERSION=X.X.X
$ yum install https://github.com/k1LoW/gh-grep/releases/download/v$GH-GREP_VERSION/gh-grep_$GH-GREP_VERSION-1_amd64.rpm
```

**apk:**

Use [apk-add-from-url](https://github.com/k1LoW/apk-add-from-url)

``` console
$ export GH-GREP_VERSION=X.X.X
$ curl -L https://git.io/apk-add-from-url | sh -s -- https://github.com/k1LoW/gh-grep/releases/download/v$GH-GREP_VERSION/gh-grep_$GH-GREP_VERSION-1_amd64.apk
```

**homebrew tap:**

```console
$ brew install k1LoW/tap/gh-grep
```

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
