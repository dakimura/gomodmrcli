# modpxycli

[![Build Status](https://github.com/dakimura/gomodmrcli/workflows/Run%20Tests/badge.svg?branch=main)](https://github.com/dakimura/gomodmrcli/actions?query=branch%3Amain)
[![codecov](https://codecov.io/gh/dakimura/gomodmrcli/branch/main/graph/badge.svg)](https://codecov.io/gh/dakimura/gomodmrcli)
[![Go Report Card](https://goreportcard.com/badge/github.com/dakimura/gomodmrcli)](https://goreportcard.com/report/github.com/gin-gonic/gin)
[![GoDoc](https://pkg.go.dev/badge/github.com/dakimura/gomodrmcli?status.svg)](https://pkg.go.dev/github.com/dakimura/gomodmrcli?tab=doc)
[![Join the chat at https://gitter.im/dakimura/gomodmrcli](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/dakimura/gomodmrcli?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![Sourcegraph](https://sourcegraph.com/github.com/dakimura/gomodmrcli/-/badge.svg)](https://sourcegraph.com/github.com/dakimura/gomodmrcli?badge)
[![Open Source Helpers](https://www.codetriage.com/dakimura/gomodmrcli/badges/users.svg)](https://www.codetriage.com/dakimura/gomodmrcli)
[![Release](https://img.shields.io/github/release/dakimura/gomodmrcli.svg?style=flat-square)](https://github.com/dakimura/gomodmrcli/releases)
[![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/dakimura/gomodmrcli)](https://www.tickgit.com/browse?repo=github.com/dakimura/gomodmrcli)


API client for https://index.golang.org/ and https://proxy.golang.org/

## Installation

```
$ go get -u github.com/dakimura/gomodmrcli
```
and import in your code
```
import "github.com/dakimura/gomodmrcli"
```

## Quick Start
```go
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dakimura/gomodmrcli"
)

func main() {
	defaultHttpClient := new(http.Client)
	// --- e.g. get dependencies of the module
	modulePath := "google.golang.org/protobuf"
	moduleVersion := "v1.26.0"
	proxyCli := gomodmrcli.NewProxyClient(defaultHttpClient)
	mf, _ := proxyCli.Mod(modulePath, moduleVersion, false)

	fmt.Printf("%s@%s is depending on:\n", modulePath, moduleVersion)
	for _, req := range mf.Require {
		fmt.Println(req.Syntax.Token[0])
	}

	// --- e.g. get modules recently synchronized to the official go mod proxy
	indexCli := gomodmrcli.NewIndexClient(defaultHttpClient)
	indices, _ := indexCli.Index(time.Now().Add(-24*time.Hour), 5, false)

	fmt.Println()
	fmt.Println("Recently updated modules are:")
	for _, index := range indices {
		fmt.Println(index.Path)
	}
}
```