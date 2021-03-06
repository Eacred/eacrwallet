jsonrpc/types
=============

[![Build Status](https://github.com/Eacred/eacrwallet/workflows/Build%20and%20Test/badge.svg)](https://github.com/Eacred/eacrwallet/actions)
[![ISC License](https://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/Eacred/eacrwallet/rpc/jsonrpc/types)

Package types implements concrete types for marshalling to and from the
eacrwallet JSON-RPC API.  A comprehensive suite of tests is provided to ensure
proper functionality.

The provided types are automatically registered with
[dcrjson](https://github.com/Eacred/eacrd/tree/master/dcrjson) when the package
is imported.  Although this package was primarily written for eacrwallet, it has
intentionally been designed so it can be used as a standalone package for any
projects needing to marshal to and from eacrwallet JSON-RPC requests and
responses.

Note that although it's possible to use this package directly to implement an
RPC client, it is not recommended since it is only intended as an infrastructure
package.  Instead, RPC clients should use the
[rpcclient](https://github.com/Eacred/eacrd/tree/master/rpcclient) package which
provides a full blown RPC client with many features such as automatic connection
management, websocket support, automatic notification re-registration on
reconnect, and conversion from the raw underlying RPC types (strings, floats,
ints, etc) to higher-level types with many nice and useful properties.

## Installation and Updating

```bash
$ go get -u github.com/Eacred/eacrwallet/rpc/jsonrpc/types
```

## License

Package types is licensed under the [copyfree](http://copyfree.org) ISC License.
