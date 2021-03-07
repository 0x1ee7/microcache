# microcache

HTTP memory cache with a simple interface and configurable TTL.
[![GoDoc](https://godoc.org/github.com/0x1EE7/microcache/acme?status.svg)](https://godoc.org/github.com/0x1EE7/microcache)

## Interface
- HTTP POST /<key> with the value as UTF-8 body.
- HTTP GET /<key> replies with the value as body or 404 if no such key exists.


## Installation
To install from source, just run:

```bash
go get -u github.com/0x1EE7/microcache
```

## Features

- Timeouts are handled over a single channel.

## Usage

```shellsession
$ ./microcache -http=0.0.0.0:80
2021/03/07 14:50:17 Server is starting... 0.0.0.0:80
2021/03/07 14:51:38 POST /go.mod 127.0.0.1:39020 201 curl/7.64.0
2021/03/07 14:51:46 GET /go.mod 127.0.0.1:39022 200 curl/7.64.0
2021/03/07 14:55:25 GET /go.mod/1 127.0.0.1:39024 404 curl/7.64.0

$ curl localhost/go.mod --data-binary @go.mod
$ curl localhost/go.mod
module github.com/0x1EE7/microcache

go 1.16
$ curl localhost/go.mod/1
not found
```

### Config & Environment
`CACHE_TTL` environment variable can be used to override `-ttl` file

### Help
```shellsession
$ microcache help
usage: microcache -http=localhost:8080 -ttl=30m0s
env: CACHE_TTL overrides -ttl
  -http string
        HTTP service address (default "localhost:8080")
  -ttl duration
        Cache TTL 30m, 1h etc. (default 30m0s)
```
