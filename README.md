# cidr

CLI application for basic operations with CIDR blocks based on https://github.com/open-policy-agent/opa

## Installation

Download the latest release for your platform from the
[releases page](https://github.com/rjuer/cidr/releases). Binaries are currently
available for all combinations of `[linux, darwin]` and `[amd64, arm64]`.
For `darwin` and `arm64` this is:
```sh
wget https://github.com/rjuer/cidr/releases/latest/download/cidr-darwin-arm64 -O $HOME/bin/cidr && chmod +x $HOME/bin/cidr
```

You can download a specific version as well:
```sh
wget https://github.com/rjuer/cidr/releases/download/v0.2.3/cidr-darwin-arm64 -O $HOME/bin/cidr && chmod +x $HOME/bin/cidr
```

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
