# AppStore SDK GO (Unofficial)

[![Build Status](https://travis-ci.org/Kachit/appstore-sdk-go.svg?branch=master)](https://travis-ci.org/Kachit/appstore-sdk-go)
[![codecov](https://codecov.io/gh/Kachit/appstore-sdk-go/branch/master/graph/badge.svg)](https://codecov.io/gh/Kachit/appstore-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/kachit/appstore-sdk-go)](https://goreportcard.com/report/github.com/kachit/appstore-sdk-go)
[![Release](https://img.shields.io/github/v/release/Kachit/appstore-sdk-go.svg)](https://github.com/Kachit/appstore-sdk-go/releases)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/kachit/appstore-sdk-go/blob/master/LICENSE)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go#third-party-apis) 

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
    "time"
    appstore_sdk "github.com/kachit/appstore-sdk-go"
)

func yourFuncName(){ 
    cfg := appstore_sdk.NewConfig("foo", "bar", "baz", "path/to/your/private.key")
    client := appstore_sdk.NewClientFromConfig(cfg, nil)
    
    //Build auth token
    err := client.Init()
    fmt.Println(err)

    //Build filter
    date, _ := time.Parse("2006-01-02", "2020-05-05")
    filter := &appstore_sdk.SalesReportsFilter{}
    filter.Daily().TypeSales().SubTypeSummary().Version10().SetReportDate(date)

    //Get data
    resp, err := client.SalesReports().GetReport(filter)
    if resp.IsSuccess() {
        reports := []*appstore_sdk.SalesReportSale{}
        err = resp.UnmarshalCSV(&reports)
        fmt.Println(reports[0])
    } else {
        var errorResult *appstore_sdk.ErrorResult
        _ = resp.UnmarshalError(&errorResult)
        err := errorResult.GetError()
        fmt.Println(err)
    }
}
