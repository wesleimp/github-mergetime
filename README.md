# github-mergetime

Lists the time it took a pull request to merge

## Install

### Pre-compiled binary

```sh
$ curl -sf https://gobinaries/wesleimp/github-mergetime | sh
```

### Docker

```sh
$ docker run --rm --privileged wesleimp/github-mergetime
```

### Compiling from source

**clone**

```sh
$ git clone git@github.com:wesleimp/github-mergetime.git

$ cd sfs
```

**download dependencies**

```sh
$ go mod download
```

**build**

```sh
$ go build -o github-mergetime main.go
```

**verify it works**

```sh
$ github-mergetime --help
```

## Usage

```sh
$ github-mergetime --help
NAME:
   github-mergetime - Lists the time it took a pull request to merge

USAGE:
   github-mergetime [options...] owner/repo

GLOBAL OPTIONS:
   --github-token value, -t value  Github access token (default: "") [$GITHUB_TOKEN]
   --verbose, -V                   Enable verbose mode (default: false)
   --page value                    Page number (default: 1)
   --per-page value                Number of records per page (default: 15)
   --help, -h                      show help (default: false)
   --version, -v                   print the version (default: false)
```

## Example

```sh
$ github-mergetime goreleaser/goreleaser
```

## LICENSE

[MIT](https://github.com/wesleimp/github-mergetime/blob/master/LICENSE)