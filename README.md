# cidr

CLI application for basic operations with CIDR blocks based on https://github.com/open-policy-agent/opa

## Installation

The easiest way to install `cidr` is via `go install`:

```shell
go install github.com/rjuer/cidr@latest
```

Your may also download the appropriate release for your platform from the
[releases page](https://github.com/rjuer/cidr/releases). Binaries are currently
available for all combinations of `[linux, darwin]` and `[amd64, arm64]`.

## Usage

Currently, `cidr` can:
- check if an IP address or CIDR block is contained in a particular CIDR block
- expand a CIDR block to get all its IP addresses

To check if CIDR block `192.168.0.0/16` contains the IP address `192.168.7.42`:
```shell
cidr contains 192.168.0.0/16 192.168.7.42
```

To check if CIDR block `192.168.0.0/16` contains the CIDR block `192.168.0.0/24`:
```shell
cidr contains 192.168.0.0/16 192.168.0.0/24
```

To get all IP addressed within the CIDR block `192.168.0.0/24`:
```shell
cidr expand 192.168.0.0/30
```
