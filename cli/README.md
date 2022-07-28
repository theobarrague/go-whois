# go-whois-cli

go-whois-cli is a simple Go module for domain and ip whois information query.

## Overview

All of domain, IP include IPv4 and IPv6, ASN are supported.

## Build

```shell
go build whois.go
```

## Usage

### whois query for domain

```shell
whois iana.org
```

### whois query for IPv6

```shell
whois 2001:500:88:200::8
```

### whois query for IPv4

```shell
whois 192.0.43.8
```

### whois query for ASN

```shell
whois AS13335
```

## License

Copyright © 2022, [Théo BARRAGUÉ](https://www.github.com/theobarrague)

Licensed under the MIT License 2.0
