## [v1.1.1](https://github.com/k1LoW/gh-grep/compare/v1.1.0...v1.1.1) - 2023-07-04
- Pin the API for getting release information to `api.github.com` by @k1LoW in https://github.com/k1LoW/gh-grep/pull/38
- Fix omissions by @k1LoW in https://github.com/k1LoW/gh-grep/pull/39
- Setup tagpr by @k1LoW in https://github.com/k1LoW/gh-grep/pull/40

## [v1.1.0](https://github.com/k1LoW/gh-grep/compare/v1.0.1...v1.1.0) (2023-03-09)

* Update go and pkgs [#36](https://github.com/k1LoW/gh-grep/pull/36) ([k1LoW](https://github.com/k1LoW))
* Bump golang.org/x/net from 0.5.0 to 0.7.0 [#34](https://github.com/k1LoW/gh-grep/pull/34) ([dependabot[bot]](https://github.com/apps/dependabot))
* Bump golang.org/x/crypto from 0.0.0-20210817164053-32db794688a5 to 0.1.0 [#33](https://github.com/k1LoW/gh-grep/pull/33) ([dependabot[bot]](https://github.com/apps/dependabot))

## [v1.0.1](https://github.com/k1LoW/gh-grep/compare/v1.0.0...v1.0.1) (2023-02-12)

* Fix handling no patterns error [#32](https://github.com/k1LoW/gh-grep/pull/32) ([k1LoW](https://github.com/k1LoW))

## [v1.0.0](https://github.com/k1LoW/gh-grep/compare/v0.11.1...v1.0.0) (2023-02-12)

* Bump up go-github-client (and mote) version [#31](https://github.com/k1LoW/gh-grep/pull/31) ([k1LoW](https://github.com/k1LoW))
* Fix handling args [#30](https://github.com/k1LoW/gh-grep/pull/30) ([k1LoW](https://github.com/k1LoW))
* Use `gh release` for detecting latest tag [#29](https://github.com/k1LoW/gh-grep/pull/29) ([k1LoW](https://github.com/k1LoW))

## [v0.11.1](https://github.com/k1LoW/gh-grep/compare/v0.11.0...v0.11.1) (2023-02-03)

* Update pkgs [#28](https://github.com/k1LoW/gh-grep/pull/28) ([k1LoW](https://github.com/k1LoW))
* FIX README [#26](https://github.com/k1LoW/gh-grep/pull/26) ([pyama86](https://github.com/pyama86))
* git.io deprecation [#25](https://github.com/k1LoW/gh-grep/pull/25) ([k1LoW](https://github.com/k1LoW))

## [v0.11.0](https://github.com/k1LoW/gh-grep/compare/v0.10.0...v0.11.0) (2022-03-29)

* Use bufio.NewReader instead of bufio.NewScanner [#24](https://github.com/k1LoW/gh-grep/pull/24) ([k1LoW](https://github.com/k1LoW))

## [v0.10.0](https://github.com/k1LoW/gh-grep/compare/v0.9.0...v0.10.0) (2021-11-16)

* Update packages and Go [#22](https://github.com/k1LoW/gh-grep/pull/22) ([k1LoW](https://github.com/k1LoW))

## [v0.9.0](https://github.com/k1LoW/gh-grep/compare/v0.8.0...v0.9.0) (2021-11-07)

* Add options `--branch` and `--tag` [#20](https://github.com/k1LoW/gh-grep/pull/20) ([k1LoW](https://github.com/k1LoW))

## [v0.8.0](https://github.com/k1LoW/gh-grep/compare/v0.7.0...v0.8.0) (2021-11-06)

* Use k1LoW/ghfs instead of johejo/ghfs [#19](https://github.com/k1LoW/gh-grep/pull/19) ([k1LoW](https://github.com/k1LoW))
* README: add MacPorts install info [#18](https://github.com/k1LoW/gh-grep/pull/18) ([herbygillot](https://github.com/herbygillot))

## [v0.7.0](https://github.com/k1LoW/gh-grep/compare/v0.6.0...v0.7.0) (2021-11-03)

* Add option `--only-matching` [#17](https://github.com/k1LoW/gh-grep/pull/17) ([k1LoW](https://github.com/k1LoW))
* Add option `--count` [#16](https://github.com/k1LoW/gh-grep/pull/16) ([k1LoW](https://github.com/k1LoW))

## [v0.6.0](https://github.com/k1LoW/gh-grep/compare/v0.5.0...v0.6.0) (2021-11-02)

* Fix script [#14](https://github.com/k1LoW/gh-grep/pull/14) ([mattn](https://github.com/mattn))
* Support colors on Windows [#13](https://github.com/k1LoW/gh-grep/pull/13) ([mattn](https://github.com/mattn))
* Add option `--url` [#15](https://github.com/k1LoW/gh-grep/pull/15) ([k1LoW](https://github.com/k1LoW))

## [v0.5.0](https://github.com/k1LoW/gh-grep/compare/v0.4.0...v0.5.0) (2021-11-01)

* Add options `--name-only` `--repo-only` [#12](https://github.com/k1LoW/gh-grep/pull/12) ([k1LoW](https://github.com/k1LoW))
* Add option `--ignore-case` [#11](https://github.com/k1LoW/gh-grep/pull/11) ([k1LoW](https://github.com/k1LoW))

## [v0.4.0](https://github.com/k1LoW/gh-grep/compare/v0.3.0...v0.4.0) (2021-11-01)

* Add option `-e` [#10](https://github.com/k1LoW/gh-grep/pull/10) ([k1LoW](https://github.com/k1LoW))
* Add option `--line-number` [#9](https://github.com/k1LoW/gh-grep/pull/9) ([k1LoW](https://github.com/k1LoW))

## [v0.3.0](https://github.com/k1LoW/gh-grep/compare/v0.2.2...v0.3.0) (2021-11-01)

* Color matched word [#8](https://github.com/k1LoW/gh-grep/pull/8) ([k1LoW](https://github.com/k1LoW))

## [v0.2.2](https://github.com/k1LoW/gh-grep/compare/v0.2.1...v0.2.2) (2021-11-01)

* Color the output delimiter. [#7](https://github.com/k1LoW/gh-grep/pull/7) ([k1LoW](https://github.com/k1LoW))

## [v0.2.1](https://github.com/k1LoW/gh-grep/compare/v0.2.0...v0.2.1) (2021-11-01)


## [v0.2.0](https://github.com/k1LoW/gh-grep/compare/v0.1.1...v0.2.0) (2021-11-01)

* Fix list repositories when org [#6](https://github.com/k1LoW/gh-grep/pull/6) ([k1LoW](https://github.com/k1LoW))
* Add debug log [#5](https://github.com/k1LoW/gh-grep/pull/5) ([k1LoW](https://github.com/k1LoW))

## [v0.1.1](https://github.com/k1LoW/gh-grep/compare/v0.1.0...v0.1.1) (2021-11-01)

* Fix handling of environment variables. [#4](https://github.com/k1LoW/gh-grep/pull/4) ([k1LoW](https://github.com/k1LoW))

## [v0.1.0](https://github.com/k1LoW/gh-grep/compare/v0.0.2...v0.1.0) (2021-11-01)


## [v0.0.2](https://github.com/k1LoW/gh-grep/compare/v0.0.1...v0.0.2) (2021-11-01)

* Add gh-extension setting [#3](https://github.com/k1LoW/gh-grep/pull/3) ([k1LoW](https://github.com/k1LoW))

## [v0.0.1](https://github.com/k1LoW/gh-grep/compare/a30357888af0...v0.0.1) (2021-10-31)

* Add release setting [#2](https://github.com/k1LoW/gh-grep/pull/2) ([k1LoW](https://github.com/k1LoW))
* Support `GH_*` tokens [#1](https://github.com/k1LoW/gh-grep/pull/1) ([k1LoW](https://github.com/k1LoW))
