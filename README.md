# cidr

CLI application for basic operations with CIDR blocks based on https://github.com/open-policy-agent/opa

## Installation

The easiest way to install `cidr` is via `go install`:

```shell
go install github.com/rjuer/cidr@latest
```

Your may also download the appropriate release for your platform from the
[releases page](https://github.com/rjuer/cidr/releases).

Binaries are currently available for all combinations of `[linux, darwin]` and `[amd64, arm64]`.

## Usage

Currently, `cidr` can:
- check if an IP address or CIDR block is contained in a particular CIDR block
- expand a CIDR block to get all its IP addresses

``` shell
cidr contains $CIDR $IP
cidr contains $CIDR $TESTCIDR
cidr expand $CIDR

cidr contains 192.168.0.0/16 192.168.7.42
cidr contains 192.168.0.0/16 192.168.0.0/24
cidr expand 192.168.0.0/24
```
