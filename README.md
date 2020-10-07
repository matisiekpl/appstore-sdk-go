# AppStore SDK GO (Unofficial)

[![Build Status](https://travis-ci.org/Kachit/appstore-sdk-go.svg?branch=master)](https://travis-ci.org/Kachit/appstore-sdk-go)
[![codecov](https://codecov.io/gh/Kachit/appstore-sdk-go/branch/master/graph/badge.svg)](https://codecov.io/gh/Kachit/appstore-sdk-go)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/kachit/appstore-sdk-go/blob/master/LICENSE)

## Description
Unofficial Golang SDK for AppStore Connect API

## API documentation
https://developer.apple.com/documentation/appstoreconnectapi/download_finance_reports

https://developer.apple.com/documentation/appstoreconnectapi/download_sales_and_trends_reports

## Download
```shell
go get -u github.com/kachit/appstore-sdk-go
```

## Usage
```go
package main

import (
    "fmt"
    "net/http"
    appstore_sdk "github.com/kachit/appstore-sdk-go"
)

func yourFuncName(){ 
    cfg := appstore_sdk.NewConfig()

    client := appstore_sdk.NewClient(cfg, &http.Client{})

    fmt.Print(client)
}
