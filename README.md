<img align="right" alt="usestdlibvars" src="./assets/uslv_logo.svg">

# usestdlibvars

A linter that detect the possibility to use variables/constants from the Go standard library.

## Install

### `go install`

```bash
go install github.com/sashamelentyev/usestdlibvars@latest
```

### `golangci-lint`

`usestdlibvars` is already integrated with
[golangci-lint](https://github.com/golangci/golangci-lint).

## Usage

### Binary

```console
$ usestdlibvars -h
usestdlibvars: A linter that detect the possibility to use variables/constants from the Go standard library.

Usage: usestdlibvars [-flag] [package]


Flags:
  -V    print version and exit
  -all
        no effect (deprecated)
  -c int
        display offending line with this many lines of context (default -1)
  -constant-kind
        suggest the use of constant.Kind.String()
  -cpuprofile string
        write CPU profile to this file
  -crypto-hash
        suggest the use of crypto.Hash.String()
  -debug string
        debug flags, any subset of "fpstv"
  -fix
        apply all suggested fixes
  -flags
        print analyzer flags in JSON
  -http-method
        suggest the use of http.MethodXX (default true)
  -http-status-code
        suggest the use of http.StatusXX (default true)
  -json
        emit JSON output
  -memprofile string
        write memory profile to this file
  -os-dev-null
        suggest the use of os.DevNull
  -rpc-default-path
        suggest the use of rpc.DefaultXXPath
  -source
        no effect (deprecated)
  -sql-isolation-level
        suggest the use of sql.LevelXX.String()
  -tags string
        no effect (deprecated)
  -test
        indicates whether test files should be analyzed, too (default true)
  -time-layout
        suggest the use of time.Layout
  -time-month
        suggest the use of time.Month.String()
  -time-weekday
        suggest the use of time.Weekday.String()
  -tls-signature-scheme
        suggest the use of tls.SignatureScheme.String()
  -trace string
        write trace log to this file
  -v    no effect (deprecated)
```

### `golangci-lint`

```console
golangci-lint run --disable-all --enable usestdlibvars
```

## Examples

```go
package response

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// JSON marshals v to JSON, automatically escaping HTML,
// setting the Content-Type header as "application/json; charset=utf-8",
// sends an HTTP response header with the provided statusCode and
// writes the marshaled v as bytes to the connection as part of an HTTP reply.
func JSON(w http.ResponseWriter, statusCode int, v any) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	if _, err := w.Write(buf.Bytes()); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
```

```bash
usestdlibvars ./...
```

```console
response.go:18:30: "500" can be replaced by http.StatusInternalServerError
response.go:24:30: "500" can be replaced by http.StatusInternalServerError
```

## Sponsors

[<img src="https://evrone.com/logo/evrone-sponsored-logo.png">](https://evrone.com/?utm_source=usestdlibvars)
