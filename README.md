# go-whois

go-whois is a simple Go module for domain and ip whois information query.

## Overview

All of domain, IP include IPv4 and IPv6, ASN are supported.

You can build the [whois client tool](cli).

Or you can do development by using this golang module as below.

## Installation

```shell
go get -u github.com/theobarrague/go-whois
```

## Importing

```go
import (
    "github.com/theobarrague/go-whois"
)
```

## Example

### whois query for domain

```go
result, err := whois.Whois("iana.org")
if err == nil {
    fmt.Println(result)
}
```

### whois query for IPv6

```go
result, err := whois.Whois("2001:500:88:200::8")
if err == nil {
    fmt.Println(result)
}
```

### whois query for IPv4

```go
result, err := whois.Whois("192.0.43.8")
if err == nil {
    fmt.Println(result)
}
```

### whois query for ASN

```go
result, err := whois.Whois("AS13335")
if err == nil {
    fmt.Println(result)
}
```

## License

Copyright © 2022, [Théo BARRAGUÉ](https://www.github.com/theobarrague)

Licensed under the MIT License
