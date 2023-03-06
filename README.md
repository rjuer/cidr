# cidr

CLI application for basic operations with CIDR blocks based on https://github.com/open-policy-agent/opa

## Installation

The binary is currently not available via a package manager.

For now, you can simply download the binary to you local machine, e.g. using

``` shell
wget https://github.com/rjuer/cidr/releases/download/v0.0.1/cidr && chmod +x cidr
```

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
